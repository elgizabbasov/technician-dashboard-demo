import Vue from 'vue'
import Vuex from 'vuex'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import Sidebar from './components/Sidebar.vue'
import TemperatureChart from './components/TemperatureChart.vue'
import SetupWiz from './components/SetupWizard.vue'
import StatCard1 from './components/StatCard1.vue'
import StatCard2 from './components/StatCard2.vue'
import SelectDevice from './components/SelectDevice.vue'
import SensorHealth from './components/SensorHealth.vue'
import TechnicianForm from './components/TechnicianForm.vue'

import msal from 'vue-msal'
import Vuetify from 'vuetify/lib'
//import * as msal from 'vue-msal';

Vue.config.productionTip = false;

// Vue.use(msal, {
//  auth: {
//    clientId: ""
//  }
// });

Vue.use(Vuex)
Vue.component("sidebar-component", Sidebar)
Vue.component("temperaturechart-component", TemperatureChart)
Vue.component("setup-component", SetupWiz)
Vue.component("statcard1-component", StatCard1)
Vue.component("statcard2-component", StatCard2)
Vue.component("select-device-component", SelectDevice)
Vue.component("sensor-health-component", SensorHealth)
Vue.component("technician-form-component", TechnicianForm)


new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')