<style>
.pieCard {
  width: 10vw;
  height: 10vh;
  background: rgba(250, 251, 252, 0.2);
  display: flex;
}
</style>

<template>
  <div class="pieCard" style="padding-left: 12px; width: 300px; height: 80px;">
    <v-card class="mx-auto" :fit="true" min-width="300">
      <v-card-text>
        <div class="text">
          <span style="font-size: 18px; color: black;">
            Days since last checkup: {{ daysSinceCheckup ? daysSinceCheckup : 'No previous checkup data found!' }}
          </span>
        </div>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      serviceID: null,
      lastCheckup: null,
    };
  },
  async mounted() {
    await this.setup();
  },
  methods: {
    async setup() {
      await this.fetchServiceId();
      await this.fetchLastCheckup();
    },
    async fetchServiceId() {
      try {
        const response = await axios.get(this.$store.getters.url + 
          "services/" +
          "deviceid?deviceid=" + 
          this.$store.getters.selectedDevice
        );
        const data = response.data.data;
        
        if (data) {
          this.serviceID = data.id;
        }
      } catch (error) {
        console.error("Service wasnt found for the device", error);
      }
    },
    async fetchLastCheckup() {
      if (!this.serviceID) {
        return;
      }
      try {
        const response = await axios.get(this.$store.getters.url + 
          "servicedata/" + 
          "serviceid?serviceid=" +
          this.serviceID
        );
        const data = response.data.data;

        if (data && data.length > 0) {
          const sortedData = data.sort((a, b) => {
            const dateA = new Date(a.created);
            const dateB = new Date(b.created);
            return dateB - dateA; // sort in descending order
          });
          this.lastCheckup = sortedData[0].created;
        }
      } catch (error) {
        this.lastCheckup = null;
        console.error(error);
      }
    },
  },
  computed: {
    daysSinceCheckup() {
      if (!this.lastCheckup) {
        return '';
      }
      const now = new Date();
      const [datePart, timePart] = this.lastCheckup.split(' ');
      const [year, month, day] = datePart.split('-').map(part => parseInt(part));
      const [hour, minute, second] = timePart.split(':').map(part => parseInt(part));
      
      const lastCheckupDate = new Date(year, month - 1, day, hour, minute, second);

      const diffInMs = now - lastCheckupDate;
      const diffInDays = Math.floor(diffInMs / (1000 * 60 * 60 * 24));
      return diffInDays;
    },
  },
  watch: {
    async '$store.state.selectedDevice'(newValue, oldValue) {
      await this.setup();
    }
  }
}
</script>
