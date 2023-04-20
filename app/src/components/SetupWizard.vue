<template>
    <v-container fluid>
      <v-card>
        <v-card-title class="deep-purple white--text">
          <span class="text-h6">Setup Wizard</span>
          <v-spacer></v-spacer>
        </v-card-title>
        <v-stepper v-model="step" flat alt-labels>
        <v-stepper-header class="elevation-0">
              <v-stepper-step
              :complete="step > 1"
              step="1"
              color=#4d00fa
            >
              Setup Site
            </v-stepper-step>

            <v-divider></v-divider>

            <v-stepper-step
              :complete="step > 2"
              step="2"
              color=#4d00fa
            >
              Setup Device
            </v-stepper-step>

            <v-divider></v-divider>

            <v-stepper-step 
              :complete="step > 3"
              step="3"
              color=#4d00fa
            > Setup Sensors 
            </v-stepper-step>

            <v-divider></v-divider>

          <v-stepper-step 
            :complete="step > 3"
            step="4"
            color=#4d00fa
          > Completion 
          </v-stepper-step>
        </v-stepper-header>
        <v-divider></v-divider>
        <v-stepper-items>
          <v-stepper-content step="1">
          <v-form ref="site_form" lazy-validation>
           <!-- <v-text-field
                v-model="SiteID"
                label="Site ID"
            ></v-text-field> -->

              <v-text-field
                  v-model="SiteName"
                  label="Site Name"
                  :rules="[required, siteNameAvailable]"
                  @blur="checkSiteNameAvailability"
                  :error-messages="siteNameErrorMessage"
              >
                <template v-slot:v-messages="{ msg }">
                  <span v-if="msg.length">{{ msg[0] }}</span>
                </template>
              </v-text-field> 

              <v-text-field
                  v-model="Organization"
                  label="Organization"
                  :rules="[required]"
              ></v-text-field> 

              <v-text-field
                  v-model="Address"
                  label="Address"
                  :rules="[required]"
              ></v-text-field> 

              <v-text-field
                  v-model="Country"
                  label="Country"
                  :rules="[required]"
              ></v-text-field> 


              <v-text-field
                  v-model="Province"
                  label="Province"
                  :rules="[v=> !!v || 'Please fill out all fields', v=> v.length == 2 || 'Please add Province in abbreviated format, ex: TX for Texas']"
              ></v-text-field> 

              <v-text-field
                  v-model="City"
                  label="City"
                  :rules="[required]"
              ></v-text-field> 

              <v-text-field
                  v-model="PostalCode"
                  label="Postal Code"
                  :rules="[v=> !!v || 'Please fill out all fields', v=> /^[a-zA-Z]{1}[0-9]{1}[a-zA-Z]{1}(| |)[0-9]{1}[a-zA-Z]{1}[0-9]{1}$/.test(v) || 'Please enter valid postal code' ]"
              ></v-text-field> 
              </v-form>
              <v-btn
                color=#4d00fa
                class="white--text"
                type = "submit"
                @click="validate_site(2)"
              >
                Continue
              </v-btn>

            </v-stepper-content>

            <v-stepper-content step="2">
              <v-form ref="device_form" lazy-validation>
              <!-- <v-text-field
                  v-model="DeviceID"
                  label="Device ID"
              ></v-text-field>  -->

              <!-- <v-text-field
                  v-model="SiteID"
                  label="Site ID"
              ></v-text-field>  -->

              <v-text-field
                  v-model="DeviceName"
                  label="Device Name"
                  :rules="[required, deviceNameAvailable]"
                  @blur="checkDeviceNameAvailability"
                  :error-messages="deviceNameErrorMessage"
              >
                <template v-slot:v-messages="{ msg }">
                  <span v-if="msg.length">{{ msg[0] }}</span>
                </template>
              </v-text-field> 

              <v-text-field
                  v-model="Version"
                  label="Version"
                  :rules="[required]"
              ></v-text-field> 

              <v-text-field
                  v-model="SerialNumber"
                  label="Serial Number"
                  :rules="[required]"
              ></v-text-field> 
              </v-form>
              <v-btn
                color=#4d00fa
                class="white--text"
                @click="validate_device(3)"
                type="submit"
              >
                Continue
              </v-btn>
              <v-btn
                text
                @click="step = 1"
              >
                Back
              </v-btn>
            </v-stepper-content>

            <v-stepper-content step="3">
                <v-form ref="sensor_form" lazy-validation>
                <!-- <v-text-field
                      v-model="SensorID"
                      label="Sensor ID"
                  ></v-text-field>  -->

                  <!-- <v-text-field
                      v-model="DeviceID"
                      label="Device ID"
                  ></v-text-field>  -->

                  <v-text-field
                      v-model="SensorName"
                      label="Sensor Name"
                      :rules="[required, sensorNameAvailable]"
                      @blur="checkSensorNameAvailability"
                      :error-messages="sensorNameErrorMessage"
                  >
                    <template v-slot:v-messages="{ msg }">
                      <span v-if="msg.length">{{ msg[0] }}</span>
                    </template>
                  </v-text-field> 

                  <v-text-field
                      v-model="Unit"
                      label="Unit"
                      :rules="[required]"
                  ></v-text-field> 

                  <v-select
                      v-model="Type"
                      label="Type"
                      :items= "['Temperature', 'Pressure', 'Humidity']"
                      :rules="[required]"
                      density: compact
                  ></v-select> 
                </v-form>
              <v-btn 
                  class="white--text"
                  color=#4d00fa
                  @click="validate_sensor(4)"
                  type="submit"
              > 
                  Submit
              </v-btn>
              <v-btn 
                  text
                  @click="step = 2"
              > 
                  Back
              </v-btn>
            </v-stepper-content>

        <v-stepper-content step="4">
            <div class="text-h3 text-center" style="color:#4d00fa">
              <v-icon color=#4d00fa size="75">mdi-check-circle-outline</v-icon>
              <strong>
                  Completed!
              </strong>
            </div>
            <v-btn 
                color=#4d00fa
                class="white--text"
                @click="step = 3"
            > 
                Add New Sensor to {{ this.DeviceName }}
            </v-btn>
        </v-stepper-content>
        </v-stepper-items>
      </v-stepper>
    </v-card>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    valid: false,
    deviceCreated: false,
    siteCreated: false,
    step: 1,

    SiteName: '',
    Organization: '',
    Address: '',
    PostalCode: '',
    City: '',
    Province: '',
    Country: '',
    
    DeviceName: '',
    Version: '',
    SiteID: '',
    SerialNumber: '',

    SensorName: '',
    Unit: '',
    Type: '',
    DeviceID: '',

    siteNameAvailable: true,
    deviceNameAvailable: true,
    sensorNameAvailable: true,
    siteNameErrorMessage: '',
    deviceNameErrorMessage: '',
    sensorNameErrorMessage: '',
    defaultNameErrorMessage: 'This name is already taken. Please choose a different name.',
  }),

  methods: {
    async checkIfExists(name, type) {
      // Send a GET request to your backend to check if the name already exists
      try {
        const response = await axios.get(this.$store.getters.url + type + "names", {
          params: { name }
        });
        return response.data.exists;
      } catch (error) {
        console.log(error) // returned 400 which means site name not found
        return false;
      }
    },
    async checkSiteNameAvailability() {
      const exists = await this.checkIfExists(this.SiteName, "sites/");
      this.siteNameAvailable = !exists;
      
      if (!this.siteNameAvailable) {
        this.siteNameErrorMessage = this.defaultNameErrorMessage;
      } else {
        this.siteNameErrorMessage = '';
      }
    },
    async checkDeviceNameAvailability() {
      const exists = await this.checkIfExists(this.DeviceName, "devices/");
      this.deviceNameAvailable = !exists;
      
      if (!this.deviceNameAvailable) {
        this.deviceNameErrorMessage = this.defaultNameErrorMessage;
      } else {
        this.deviceNameErrorMessage = '';
      }
    },
    async checkSensorNameAvailability() {
      const exists = await this.checkIfExists(this.SensorName, "sensors/");
      this.sensorNameAvailable = !exists;
      
      if (!this.sensorNameAvailable) {
        this.sensorNameErrorMessage = this.defaultNameErrorMessage;
      } else {
        this.sensorNameErrorMessage = '';
      }
    },
    async validate_site(step) {
      await this.checkSiteNameAvailability();

      this.valid = false;
      this.deviceCreated = false;
      this.siteCreated = false;

      let valid = this.$refs.site_form.validate();
      if (valid && this.siteNameAvailable){
        this.step = step;
      }
    },
    async validate_device(step) {
      await this.checkDeviceNameAvailability();

      this.valid = false;
      this.deviceCreated = false;
      
      let valid = this.$refs.device_form.validate();
      if (valid && this.deviceNameAvailable){
        this.step = step;
      }
    },
    async validate_sensor(step) {
      await this.checkSensorNameAvailability();
      
      let valid = this.$refs.sensor_form.validate();
      if (valid && this.sensorNameAvailable){
        this.step = step;

        if (this.deviceCreated){
          this.submitSensorSetup();
        } else {
          if (this.siteCreated){
            this.submitDeviceSetup();
          } else {
            this.submitTotalSetup();
          }
        }
      }
    },
    async submitTotalSetup() {
      const response = await axios.post(this.$store.getters.url + 'sites/', {
        name: this.SiteName,
        organization: this.Organization,
        address: this.Address,
        postal: this.PostalCode,
        city: this.City,
        province: this.Province,
        country: this.Country,
      })
      .then(response => {
        this.SiteID = response.data.id;
        return axios.post(this.$store.getters.url + 'devices/', {
          name: this.DeviceName,
          version: this.Version,
          site_id: this.SiteID,
          serial_number: this.SerialNumber,
        })
      })
      .then(response => {
        this.DeviceID = response.data.id;
        return axios.post(this.$store.getters.url + 'sensors/', {
          name: this.SensorName,
          unit: this.Unit,
          type: this.Type,
          device_id: this.DeviceID,
        })
      })
      .catch(error => console.log(error))
      .finally(() => {
        this.deviceCreated = true;
        this.siteCreated = true;
      });
    },
    async submitDeviceSetup() {
      try {
        const response = await axios.post(this.$store.getters.url + 'devices/', {
          name: this.DeviceName,
          version: this.Version,
          site_id: this.SiteID,
          serial_number: this.SerialNumber,
        })
        .then(response => {
          this.DeviceID = response.data.id;
          return axios.post(this.$store.getters.url + 'sensors/', {
            name: this.SensorName,
            unit: this.Unit,
            type: this.Type,
            device_id: this.DeviceID,
          })
        })
        .finally(() => {
          this.deviceCreated = true;
        });
        console.log(response.data.message);
      } catch (error) {
        console.error(error);
      }
    },
    async submitSensorSetup() {
      try {
        const response = await axios.post(this.$store.getters.url + 'sensors/', {
          name: this.SensorName,
          unit: this.Unit,
          type: this.Type,
          device_id: this.DeviceID,
        });
        console.log(response.data.message);
      } catch (error) {
        console.error(error);
      }
    },
  }, 
  computed: {
    required() {
      return v => !!v || 'Please fill out all fields';
    }
  }
}
</script>