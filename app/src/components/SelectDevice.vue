<template>
    <div>
      <v-card color="#dee8ff"> 
        <v-card-title>
          <span class="text-h6">Select Device</span>
          <v-spacer></v-spacer>
        </v-card-title>
        <v-list>
              <v-list-item v-for="device in devices" :key="device.id" @click="updateOption(device)">
                  <v-list-item-content>
                      <v-list-item-title>{{ device.name }}</v-list-item-title>
                  </v-list-item-content>
              </v-list-item>
        </v-list>
      </v-card>
    </div>
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