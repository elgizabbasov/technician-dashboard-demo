<!-- <style>
    /* .chartMenu{
        width: 100px;
        height: 40px;
        background: grey;
        color: rbga(255, 25, 0, 0);
    }
    .chartMenu{
        padding: 10px;
        font-size: 20px;
    } */
    .chartCard{
        width: 100vw;
        height: 35vh;
        background: rgba(250, 251, 252, 0.2);
        display: flex;
    }
    .chartBox{
        width: 100%;
        padding: 0;
        border-radius: 2;
        border: solid 3px rgba(255, 26, 0, 0);
        background: white;
    }    
</style> -->

<template>
    <v-container fluid>
        <v-row>
            <v-col>
                <v-card flat>
                    <v-list-item-title class="text-h5 text-wrap"> Selected Device: {{ selectedDeviceName }} </v-list-item-title>
                    <v-list-item-subtitle class="text-wrap"> GUID: {{ selectedDeviceGUID }} </v-list-item-subtitle>
                </v-card>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <v-card>
                    <v-card-title class="deep-purple white--text">
                        <span class="text-h6">Sensor Readings</span>
                        <v-spacer></v-spacer>
                    </v-card-title>
                    <v-row>
                      <v-col>
                        <v-btn @click="hundredpointpreset"> Preset: Most Recent 100 Minutes </v-btn>
                        <v-btn @click="reset"> Reset </v-btn>
                      </v-col>
                    </v-row>
                    <v-row>
                        <v-col>
                          <v-card>
                              <v-list-item-title class="text-h5 text-center"> Pressure </v-list-item-title>
                              <div class="chartCenter" style="position: relative; margin: auto">
                                <canvas ref="pressureCanvas" v-for="(item, index) in presSensors" :key="index" :id="'PressureChart-' + index" :width="canvasWidth" :height="canvasHeight"></canvas>
                              </div>
                          </v-card>
                        </v-col>
                        <v-col>
                          <v-card>
                              <v-list-item-title class="text-h5 text-center"> Temperature </v-list-item-title>
                              <div class="chartCenter" style="position: relative; margin: auto">
                                <canvas ref="temperatureCanvas" v-for="(item, index) in tempSensors" :key="index" :id="'TemperatureChart-' + index" :width="canvasWidth" :height="canvasHeight"></canvas>
                              </div>
                          </v-card>
                        </v-col>
                        <v-col>
                          <v-card>
                              <v-list-item-title class="text-h5 text-center"> Humidity </v-list-item-title>
                              <div class="chartCenter" style="position: relative; margin: auto">
                                <canvas ref="humidityCanvas" v-for="(item, index) in humidSensors" :key="index" :id="'HumidityChart-' + index" :width="canvasWidth" :height="canvasHeight"></canvas>
                              </div>
                          </v-card>
                        </v-col>
                    </v-row>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import axios from "axios";
import store from "../store"
import Chart from "chart.js/auto";
import { VListItemTitle, VListItemSubtitle } from 'vuetify/lib'

export default {
  components: {
    VListItemTitle,
    VListItemSubtitle,
  },
  data() {
    return {
      Device: "",
      DeviceID: "",
      sensors: "",
      tempSensors: [],
      presSensors: [],
      humidSensors: [],
      //Row 1: Pressure, Row 2: Temperature, Row 3: Humidity
      sensorData: [[],[],[]],
      sensorPointDate: [[],[],[]],
      instance_pcharts: [],
      instance_tcharts: [],
      instance_hcharts: [],
      ctx_p: [],
      ctx_t: [],
      ctx_h: [],
      pressureDataArray: [],
      temperatureDataArray: [],
      humidityDataArray: [],

      defaultNumberOfPoints: 333,
      windowWidth: window.innerWidth,
      isResizing: false,
      // Not the best way - kepping track of what is window size to display charts in good heights
      isSmallScreen: window.innerWidth < 768,
      canvasWidth: 2000,
      canvasHeightForLargeWindowSize: 300,
      canvasHeight: this.isSmallScreen ? this.canvasWidth : this.canvasHeightForLargeWindowSize,
    };
  },

  created() {
    this.canvasHeight = this.isSmallScreen ? this.canvasWidth : this.canvasHeightForLargeWindowSize; // if window small, make chart size same_width X same_height
  },

  async mounted() {
    window.addEventListener('resize', this.handleResize);
    await this.setup();
  },

  beforeDestroy() {
    window.removeEventListener('resize', this.handleResize);
  },

  methods: {
    async handleResize() {
      if (this.isResizing) {
        return;
      }
      this.isResizing = true;
      if (window.innerWidth !== this.windowWidth) { // reset to default view if window size changes
        this.windowWidth = window.innerWidth;
        this.isSmallScreen = window.innerWidth < 768; // check if screen size is small according to https://mdbootstrap.com/docs/vue/layout/breakpoints/ 
        this.canvasHeight = this.isSmallScreen ? this.canvasWidth : this.canvasHeightForLargeWindowSize; // set to 2000 (canvasWidth) if small so charts are like squares
        await this.update(this.defaultNumberOfPoints);
      }
      this.isResizing = false;
    },
    async getSensors() {
      this.sensors = "";
      this.tempSensors = [];
      this.presSensors = [];
      this.humidSensors = [];
      try {
        const response = await axios.get(
          this.$store.getters.url + 
            "sensors/" +
            "deviceid?deviceid=" +
            this.$store.getters.selectedDevice
        );
        this.sensors = JSON.parse(JSON.stringify(response.data.data));
      } catch (error) {
        console.error(error);
      }
    },
    sortSensors() {
      for (let i = 0; i < this.sensors.length; i++) {
        if (this.sensors[i].type === "Pressure") {
          this.presSensors.push(this.sensors[i]);
        }
        if (this.sensors[i].type === "Temperature") {
          this.tempSensors.push(this.sensors[i]);
        }
        if (this.sensors[i].type === "Humidity") {
          this.humidSensors.push(this.sensors[i]);
        }
      }
    },
    async getSensorData(rows, sensorid, type) {
      try {
        const response = await axios.get(
          this.$store.getters.url +
          "sensordata/" +
          "latest?sensorid=" +
          sensorid +
          "&rows=" +
          rows.toString()
        );
        let rArray = [];
        for(let i = 0;i < response.data.data.length; i++) {
          rArray.push(parseFloat(response.data.data[i].data));
          if(type === 'pressure' && i === 0) {
            this.$store.dispatch('setPSensorZero', response.data.data[i].created)
          } else if(type === 'temperature' && i === 0) {
            this.$store.dispatch('setTSensorZero', response.data.data[i].created)
          } else if(type === 'humidity' && i === 0) {
            this.$store.dispatch('setHSensorZero', response.data.data[i].created)
          }
        }
        return rArray;
      } catch (error) {
          console.error(error);
      }
    },
    async populateSensorDataArrays(points) {
      //reset store values from previous use
      this.$store.dispatch('reset');
      //check if presSensors has sensors in the array
      if(this.presSensors.length > 0) {
        for(let i = 0;i < this.presSensors.length; i++) {
          //this pushes an array of floats into the first row of sensorData making it a 2d array of sensordata. Pres data is in the first row, temp in second, humidity in third.
          this.sensorData[0].push(await this.getSensorData(points, this.presSensors[i].id, 'pressure'));
        }
      } else {
        this.sensorData[0] = [];
      };
      //check if there are sensors in the tempSensor array
      if(this.tempSensors.length > 0) {
        for(let i = 0;i < this.tempSensors.length; i++) {
          this.sensorData[1].push(await this.getSensorData(points, this.tempSensors[i].id, 'temperature'));
        }
      } else {
        this.sensorData[1] = [];
      };
      if(this.humidSensors.length > 0) {
        for(let i = 0;i < this.humidSensors.length; i++) {
          this.sensorData[2].push(await this.getSensorData(points, this.humidSensors[i].id, 'humidity'));
        }
      } else {
        this.sensorData[2] = [];
      };
    },
    async setupSensors(points) {
      await this.getSensors();
      await this.sortSensors();
      await this.populateSensorDataArrays(points);
    },
    async setupGraphs() {
      //pressure sensor data
      for(let i = 0; i < this.presSensors.length; i++) {
        const pressure_data = {
          labels: [],
          datasets: [
            {
              label: "Pressure Reading",
              data: this.sensorData[0][i],
              fill: true,
              borderColor: "rgb(255, 152, 0)",
              backgroundColor: "rgba(255,197,111,0.5)",
              tension: 0.3,
            },
          ],
        };
        this.pressureDataArray.push(pressure_data);
      }
      //create multiple graph id's for instancing them individually
      for(let i = 0; i < this.presSensors.length; i++) {
        this.ctx_p.push(document.getElementById("PressureChart-" + i.toString()));
      }
      // temperature sensor data
      for(let i = 0; i < this.tempSensors.length; i++) {
          const temperature_data = {
            labels: [],
            datasets: [
              {
              label: "Temperature Reading",
              data: this.sensorData[1][0],
              fill: true,
              borderColor: "rgb(244, 67, 54)",
              backgroundColor: "rgba(241, 110, 103, 0.5)",
              tension: 0.3,
              },
            ],
          };
          this.temperatureDataArray.push(temperature_data);
        }
        //create multiple graph id's for instancing them individually
        for(let i = 0; i < this.tempSensors.length; i++) {
          this.ctx_t.push(document.getElementById("TemperatureChart-" + i.toString()));
        }
      // humidity sensor data
      for(let i = 0; i < this.humidSensors.length; i++) {
          const humidity_data = {
            labels: [],
            datasets: [
              {
              label: "Humidity Reading",
              data: this.sensorData[2][0],
              fill: true,
              borderColor: "rgb(76, 175, 80)",
              backgroundColor: "rgba(154, 210, 156, 0.5)",
              tension: 0.3,
              },
            ],
          };
          this.humidityDataArray.push(humidity_data);
        }
        //create multiple graph id's for instancing them individually
        for(let i = 0; i < this.humidSensors.length; i++) {
          this.ctx_h.push(document.getElementById("HumidityChart-" + i.toString()));
        }
      //Multiple pressure chart instances initialized into instance_pcharts
      for(let i = 0; i < this.presSensors.length; i++){
        const pChart = new Chart(this.ctx_p[i], {
          type: "line",
          data: this.pressureDataArray[i],
          options: {
            layout: {
              padding: 10
            },
            scales: {
              x: {
                ticks: {
                  callback: function(value, index, values) {
                    if (index === 0 || index === Math.floor(values.length/2) || index === values.length-1) {
                      return value;
                    }
                    return '';
                  }
                }
              }
            }
          }
        })
        pChart.data.labels = Array.from({ length: this.sensorData[0][i].length }, (_, i) => (i + 1).toString());
        pChart.update();
        this.instance_pcharts.push(pChart);
      };
      
      // temperature render init block
      for(let i = 0; i < this.tempSensors.length; i++){
        const tChart = new Chart(this.ctx_t[i], {
          type: "line",
          data: this.temperatureDataArray[i],
          options: {
            layout: {
              padding: 10
            },
          scales: {
            x: {
              ticks: {
                callback: function(value, index, values) {
                  if (index === 0 || index === Math.floor(values.length/2) || index === values.length-1) {
                    return value;
                  }
                  return '';
                }
              }
            }
          }
        }
        })
        tChart.data.labels = Array.from({ length: this.sensorData[1][i].length }, (_, i) => (i + 1).toString());
        tChart.update();
        this.instance_tcharts.push(tChart);
      };
      // Humidity render init block
      for(let i = 0; i < this.humidSensors.length; i++){
        const hChart = new Chart(this.ctx_h[i], {
          type: "line",
          data: this.humidityDataArray[i],
          options: {
            layout: {
              padding: 10
            },
          scales: {
            x: {
              ticks: {
                callback: function(value, index, values) {
                  if (index === 0 || index === Math.floor(values.length/2) || index === values.length-1) {
                    return value;
                  }
                  return '';
                }
              }
            }
          }
        }
        })
        hChart.data.labels = Array.from({ length: this.sensorData[2][i].length }, (_, i) => (i + 1).toString());
        hChart.update();
        this.instance_hcharts.push(hChart);
      };
    },
    async setup() {
      this.sensorData = [[],[],[]];
      await this.setupSensors(this.defaultNumberOfPoints); // sample number of points to be loaded by default 
      await this.setupGraphs();
      this.$store.dispatch('setPSensor', this.presSensors.length)
      this.$store.dispatch('setTSensor', this.tempSensors.length)
      this.$store.dispatch('setHSensor', this.humidSensors.length)
    },
    async update(points) {
      this.sensorData = [[],[],[]];
      this.$store.dispatch('reset');
      await this.setupSensors(points);
      this.$store.dispatch('setPSensor', this.presSensors.length)
      this.$store.dispatch('setTSensor', this.tempSensors.length)
      this.$store.dispatch('setHSensor', this.humidSensors.length)
      this.instance_pcharts = [];
      this.instance_tcharts= [];
      this.instance_hcharts = [];
      this.ctx_p = [];
      this.ctx_t = [];
      this.ctx_h = [];
      this.pressureDataArray = [];
      this.temperatureDataArray = [];
      this.humidityDataArray = [];
      await this.setupGraphs();
    },
    async hundredpointpreset() {
      await this.update(100);
    }, 
    async reset() {
      await this.update(this.defaultNumberOfPoints); // reset to the default number set in setup
    }
  },

  computed: {
    selectedDeviceName() {
      this.Device = this.$store.getters.selectedDeviceName;
      return this.$store.getters.selectedDeviceName;
    },
    selectedDeviceGUID() {
      this.DeviceID = this.$store.getters.selectedDevice; // will return the GUID
      return this.$store.getters.selectedDevice;
    },
  },

  watch: {
    async '$store.state.selectedDevice'(newValue, oldValue) {
      await this.update(this.defaultNumberOfPoints);
    }
  }
};
</script>

<style>
.chartCenter {
  display: flex;
  flex-wrap: wrap;
}
</style>