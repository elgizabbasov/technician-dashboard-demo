<template>
  <v-app>
    <Spinner v-if="not_authenticated" :start="not_authenticated" />
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
import Store from "./store";
import HomePage from "./views/Dashboard.vue";
import Sidebar from "./components/Sidebar.vue";
import TemperatureChart from "./components/TemperatureChart.vue";
import StatCard1 from "./components/StatCard1.vue";
import SensorHealth from './components/SensorHealth.vue';
import Spinner from '@/components/Spinner.vue';
//import { msalMixin } from 'vue-msal';
//import StatCard2 from './components/StatCard2.vue'


// import TechnicianForm from './views/TechnicianForm.vue'

export default {
  // :start="spinnerStart" @toggle-spinner="spinnerStart = $event
  name: 'App',
  components: {
    'HomePage': HomePage,
    'Sidebar' : Sidebar,
    'StatCard1' : StatCard1,
    'SensorHealth' : SensorHealth,
    'TemperatureChart': TemperatureChart,
    'Spinner' : Spinner
  },
  data() {
    return {
      not_authenticated: true
    };
  },
  created() {
    // check if showComponent value is stored in local storage
    const not_auth = localStorage.getItem('not_authenticated');
    // if storedValue is not null, parse it as a boolean and assign it to showComponent
    if (not_auth !== null) {
      this.not_authenticated = JSON.parse(not_auth);
    }
    
    if (!this.$msal.isAuthenticated()) {
      this.$msal.signIn();
    } else {
      this.not_authenticated = false; // hide the component
      // store the showComponent value in local storage as a string
      localStorage.setItem('not_authenticated', JSON.stringify(this.not_authenticated));
    }
  },
  // mounted() {
  //   const not_auth = localStorage.getItem('not_authenticated');
  //   // if storedValue is not null, parse it as a boolean and assign it to showComponent
  //   if (not_auth !== null) {
  //     this.not_authenticated = JSON.parse(not_auth);
  //   }
  // }
};
</script>

<style>
  @import url('https://fonts.googleapis.com/css?family=Manrope');
  #app {
    font-family: 'Manrope', sans-serif;
    background-color: #4D00FA;
    
  }
</style>