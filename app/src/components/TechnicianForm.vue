<template>
    <v-container fluid>
        <v-dialog class="mx-auto" max-width="360" v-model="show">
            <v-card v-if="serviceIDNotFound" class="mx-auto" max-width="360">
                <v-card-title class="justify-center">
                    <v-icon color="red" size="75">mdi-alert-circle-outline</v-icon>
                </v-card-title>
                <v-card-text>
                    <div class="text-h3 text-center" style="color: red">
                        <strong>
                            Error Submitting, Service Not Found!
                        </strong>
                    </div>
                </v-card-text>
            </v-card>
            <v-card v-else class="mx-auto" max-width="360">
                <v-card-title class="justify-center">
                    <v-icon color=#4d00fa size="75">mdi-check-circle-outline</v-icon>
                </v-card-title>
                <v-card-text>
                    <div class="text-h3 text-center" style="color:#4d00fa">
                        <strong>
                            Submitted!
                        </strong>
                    </div>
                </v-card-text>
            </v-card>
        </v-dialog>
        <v-card>
            <v-card-title class="deep-purple white--text">
                <span class="text-h6">Technician Form</span>
                <v-spacer></v-spacer>
            </v-card-title>
            <v-sheet>
                <v-form @submit.prevent ref="form" style="padding: 24px 24px 16px 24px;">
                    <v-text-field
                        v-model="form.TechnicianName"
                        label="Technician Name"
                        :rules="[required]"
                    ></v-text-field>

                    <v-text-field
                        v-model="form.Company"
                        label="Company"
                        :rules="[required]"
                    ></v-text-field>

                    <v-text-field
                        v-model="form.Address"
                        label="Address"
                        :rules="[required]"
                    ></v-text-field>

                    <v-text-field
                        v-model="form.Date"
                        label="Date"
                        placeholder="MM/DD/YYYY"
                        :rules="[v=> !!v || 'Please fill out all fields', v=> /^(0?[1-9]|1[012])\/(0?[1-9]|[12][0-9]|3[01])\/\d{4}$/.test(v) || 'Please enter date in format MM/DD/YYYY' ]"
                    ></v-text-field>

                    <v-text-field
                        v-model="form.DateKOH"
                        label="Date KOH Added"
                        placeholder="MM/DD/YYYY"
                        :rules="[v=> !!v || 'Please fill out all fields', v=> /^(0?[1-9]|1[012])\/(0?[1-9]|[12][0-9]|3[01])\/\d{4}$/.test(v) || 'Please enter date in format MM/DD/YYYY' ]"
                    ></v-text-field>

                    <v-text-field
                        v-model="form.BatchNumber"
                        label="Batch Number"
                        :rules="[required]"
                    ></v-text-field>

                    <v-text-field
                        v-model="form.NumberBags"
                        label="Number Of 25kg Bags"
                        :rules="[v=> !!v || 'Please fill out all fields', v=> /^-?\d+$/.test(v) || 'Please enter a number']"
                    ></v-text-field>

                    <v-text-field
                        v-model="form.PH"
                        label="pH"
                        :rules="[v=> !!v || 'Please fill out all fields', v=> /^-?\d*(\.\d+)?$/.test(v) || 'Please enter a decimal value']"
                    ></v-text-field>

                    <v-checkbox
                        v-model="form.Checkbox1"
                        label="Shut off power to unit"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox2"
                        label="Check pH of pearl ash record results"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox3"
                        label="Vacuum out all pearl ash from chamber"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox4"
                        label="Reload chamber with KOH and record batch number"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox5"
                        label="Ensure pails are weighed, labeled with weight & location, and sealed"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox6"
                        label="Pull filter from inside unit and rinse & leave out to dry, replace filter with a clean dry one"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox7"
                        label="Close up unit and turn power back on to unit"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox8"
                        label="Test fire appliance that the unit is attached to"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox9"
                        label="Check room for leaks"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox10"
                        label="Check pump for leaks and noises"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox11"
                        label="Wipe down all appliances in mechanical room"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-checkbox
                        v-model="form.Checkbox12"
                        label="Clean floor in mechanical room"
                        color=#4d00fa
                        :rules="[required]"
                    ></v-checkbox>

                    <v-file-input
                        multiple
                        :rules="[v => v.length > 0 || 'Please fill out all fields']"
                        v-model="images.pearlash"
                        prepend-icon="mdi-camera"
                        placeholder="Image of Pearl Ash in Chamber"
                        @change="onPearlashChange"
                    ></v-file-input>

                    <v-file-input
                        multiple
                        :rules="[v => v.length > 0 || 'Please fill out all fields']"
                        v-model="images.mechanicalroom"
                        prepend-icon="mdi-camera"
                        placeholder="Image of Mechanical Room"
                        @change="onMechanicalroomChange"
                    ></v-file-input>

                    <v-file-input
                        multiple
                        :rules="[v => v.length > 0 || 'Please fill out all fields']"
                        v-model="images.pails"
                        prepend-icon="mdi-camera"
                        placeholder="Image of Pails with Labels"
                        @change="onPailsChange"
                    ></v-file-input>

                    <v-btn
                        type="submit"
                        class="white--text"
                        @click="validate()"
                        color=#4d00fa
                    >
                        Submit
                    </v-btn>
                </v-form>
            </v-sheet>    
        </v-card>
        
    </v-container>
</template>

<script>
export default {
  data: () => ({
    valid: false,
    show: false,
    form: {
      TechnicianName: "",
      Company: "",
      Address: "",
      Date: "",
      DateKOH: "",
      BatchNumber: "",
      NumberBags: "",
      PH: "",
      Checkbox1: false,
      Checkbox2: false,
      Checkbox3: false,
      Checkbox4: false,
      Checkbox5: false,
      Checkbox6: false,
      Checkbox7: false,
      Checkbox8: false,
      Checkbox9: false,
      Checkbox10: false,
      Checkbox11: false,
      Checkbox12: false,
    },
    images: {
      pearlash: [],
      mechanicalroom: [],
      pails: [],
    },
    ServiceID: "",
    serviceIDNotFound: false
  }),
  methods: {
    // handle image uploads
    onPearlashChange(files) {
      this.images['pearlash'] = files;
    },
    onMechanicalroomChange(files) {
      this.images['mechanicalroom'] = files;
    },
    onPailsChange(files) {
      this.images['pails'] = files;
    },
    // check fields valid  
    async validate() {
        this.valid = false;
        let valid = this.$refs.form.validate();
        if (valid) {
            try {
                await this.submitForm();
                this.show = true;
                setTimeout(() => {
                    this.show = false;
                }, 2000);
            } catch (error) {
                // handle any errors that occurred during submit
                console.error(error);
            } finally {
                // always reset the form after submit, whether it succeeded or failed
                this.$refs.form.reset();
            }
        }
    },
    async submitForm() {
      try {
        const dataObject = {
            type: 'report',
            data: {} // data to be saved under service_data: data
        };
        
        // loop over the form data and append to the FormData object
        for (let key in this.form) {
            dataObject.data[key] = this.form[key];
        }
        // object holding all form data and images
        const serviceData = {
            data: JSON.stringify(dataObject),
            pearlash: [],
            mechanicalroom: [],
            pails: [],
            service_id: this.ServiceID
        };
        try {
            // send the initial request to check if service exists
            const response = await axios.get(
                this.$store.getters.url +
                "services/" +
                "deviceid?deviceid=" +
                this.$store.getters.selectedDevice
            );
            if (response && response.status === 200 && response.data && response.data.data && response.data.data.id) {
                this.ServiceID = response.data.data.id;
                serviceData.service_id = this.ServiceID;
            } else {
                this.serviceIDNotFound = true;
                console.log("ServiceID not found for given DeviceID, submit failed");
                return;
            }
        } catch (error) {
            this.serviceIDNotFound = true;
            console.log("ServiceID not found for given DeviceID, submit failed");
            return;
        }
        // loop over the images and upload them to the file server
        for (let key in this.images) {
            if (this.images[key] && this.images[key].length > 0) {
                let files = this.images[key];
                
                // Loop over each file and upload it to the server
                for (let i = 0; i < files.length; i++) {
                    const image = files[i];
                    const imageURL = await this.uploadImageToServer(image);
                    
                    // add the image URL to the service data object
                    serviceData[key].push(imageURL);
                }
            }
        }
        // post the service data to the backend
        const serviceResponse = await axios.post(this.$store.getters.url + "servicedata/", serviceData);
        console.log(serviceResponse.data.message);
      } catch (error) {
        console.error(error);
      }
    },
    async uploadImageToServer(image) {
        try {
            const formData = new FormData();
            formData.append('image', image);
            const response = await axios.post(this.$store.getters.url + 
                'images/', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
                }
            );
            console.log("Image uploaded:", response.data);
            return response.data.url;
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
};
</script>