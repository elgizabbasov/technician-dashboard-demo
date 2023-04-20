<template>
  <v-container fluid>
    <v-card width="300px">
      <v-card-title class="deep-purple white--text">
        <span class="text-h6">Sensor Health</span>
        <v-spacer></v-spacer>
      </v-card-title>
      <v-list>
        <v-list-item
          v-for="sensor in sensors"
          :key="sensor.type"
        >
          <v-list-item-content class="text-h6">
            <v-list-item-title>{{ sensor.type }}</v-list-item-title>
            <v-list-item-subtitle>{{ sensor.health }}</v-list-item-subtitle>
          </v-list-item-content>
          <v-spacer></v-spacer>
          <v-list-item-avatar>
            <v-icon :color="sensor.color">mdi-circle</v-icon>
          </v-list-item-avatar>
        </v-list-item>
      </v-list>
    </v-card>
  </v-container>  
</template>

<script>
import { VList, VListItem, VListItemContent, VListItemTitle, VListItemSubtitle } from 'vuetify/lib'

export default {
  name: 'App',
  components: {
    VList,
    VListItem,
    VListItemContent,
    VListItemTitle,
    VListItemSubtitle,
  },
  data: () => ({
    sensors: [],
  }),
  
  mounted() {
    //add sensors into item list
    this.addSensors();
  },
  methods: {
    addSensors() {
      setTimeout(() => {
        //add pressure sensor checkers based on number of sensors found
        for(let i = 0; i < this.$store.getters.pSensorLength; i++) {
          const array = this.$store.getters.pSensorZero;
          const j = i + 1;
          // check whether latest created sensordata is more than 24 hr 
          if (Math.abs(Date.now() - Date.parse((array[j])) >= 86400000)) {
            this.sensors.push({type: 'Pressure ' + j.toString(), health: 'Bad', color: 'red'});
          } else if(Math.abs(Date.now() - Date.parse((array[j]))) >= 3600000) {  // yellow if more than 1 hr 
            this.sensors.push({type: 'Pressure ' + j.toString(), health: 'Needs Review', color: 'yellow'});
          } else {
            this.sensors.push({type: 'Pressure ' + j.toString(), health: 'Good', color: 'green'});
          }
        }
        //add temperature sensor checkers based on number of sensors found
        for(let i = 0; i < this.$store.getters.tSensorLength; i++) {
          const array = this.$store.getters.tSensorZero;
          const j = i + 1;
          if (Math.abs(Date.now() - Date.parse((array[j])) >= 86400000)) {
            this.sensors.push({type: 'Temperature ' + j.toString(), health: 'Bad', color: 'red'});
          } else if(Math.abs(Date.now() - Date.parse((array[j]))) >= 3600000) {
            this.sensors.push({type: 'Temperature ' + j.toString(), health: 'Needs Review', color: 'yellow'});
          } else {
            this.sensors.push({type: 'Temperature ' + j.toString(), health: 'Good', color: 'green'});
          }
        }
        //add humidity sensor checkers based on number of sensors found
        for(let i = 0; i < this.$store.getters.hSensorLength; i++) {
          const array = this.$store.getters.hSensorZero;
          const j = i + 1;
          if (Math.abs(Date.now() - Date.parse((array[j])) >= 86400000)) {
            this.sensors.push({type: 'Humidity ' + j.toString(), health: 'Bad', color: 'red'});
          } else if(Math.abs(Date.now() - Date.parse((array[j]))) >= 3600000) {
            this.sensors.push({type: 'Humidity ' + j.toString(), health: 'Needs Review', color: 'yellow'});
          } else {
            this.sensors.push({type: 'Humidity ' + j.toString(), health: 'Good', color: 'green'});
          }
        }
      }, 100) // TODO: Fix so not using constant timeout anymore
    }
  },
  watch: {
    '$store.state.selectedDevice'(newValue, oldValue) {
      this.sensors = [];
      setTimeout(() => {
        this.addSensors();
      }, 50) // TODO: Fix so not using constant timeout anymore
    }
  }
  
};
</script>