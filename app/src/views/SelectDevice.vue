<template>
    <v-container fill-height>
        <v-row>
            <v-col>
              <v-card>
              <select-device-component/>
                    <v-dialog
                        v-model="dialog"
                        persistent
                        max-width="600px"
                    >
                        <template v-slot:activator="{ on, attrs }">
                            <v-btn
                            v-bind="attrs"
                            v-on="on"
                            class="mx-auto elevation black--text text-center"
                            block
                            text
                            outlined
                            style="background-color: #dee8ff;"
                            >
                            Setup New Device
                            </v-btn>
                        </template>
                            <setup-component/>                    
                        <v-btn
                            color="red"
                            @click="dialog = false"
                        >
                            Close
                        </v-btn>
                    </v-dialog>
                  </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>


<script>
import store from '../store'
import { mapActions } from 'vuex'
  export default {
    data: () => ({
      selectedDevice: '',
      selectedDeviceName: '',
      devices: [''],
      dialog: false,
    }),

    methods: {
      ...mapActions(['store/setDevice'], ['store/setDeviceName']),
      async getDevices() {
        try {
          const response = await axios.get(this.$store.getters.url + "devices/")
          this.devices = JSON.parse(JSON.stringify(response.data.data)) // response.name_of_json_containing_devices.data
        } catch (error) {
          console.error(error)
        }
      },
      updateOption(device) {
        this.selectedDevice = device.id;
        this.selectedDeviceName = device.name;
        this.$store.dispatch('setDevice', this.selectedDevice)
        this.$store.dispatch('setDeviceName', this.selectedDeviceName)
        //check if current path is /dashboard
        if(this.$route.path !== '/dashboard') {
          this.$router.push("/dashboard")
        }
      }
    },
    async mounted() {
        await this.getDevices()
    },
  }
</script>