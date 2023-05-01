import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    url: 'https://technician-dashboard-demo.herokuapp.com/',
    selectedDevice: '',
    selectedDeviceName: '',
    pSensorLength: 0,
    tSensorLength: 0,
    hSensorLength: 0,
    pSensorZero: [null],
    tSensorZero: [null],
    hSensorZero: [null],
  },
  mutations: {
    SET_DEVICE(state, device) {
      state.selectedDevice = device
    },
    SET_DEVICENAME(state, deviceName) {
      state.selectedDeviceName = deviceName
    },
    SET_PSENSOR(state, pSensor) {
      state.pSensorLength = pSensor
    },
    SET_TSENSOR(state, tSensor) {
      state.tSensorLength = tSensor
    },
    SET_HSENSOR(state, hSensor) {
      state.hSensorLength = hSensor
    },
    APPEND_PSENSORZERO(state, pSensorPoint) {
      state.pSensorZero.push(pSensorPoint)
    },
    APPEND_TSENSORZERO(state, tSensorPoint) {
      state.tSensorZero.push(tSensorPoint)
    },
    APPEND_HSENSORZERO(state, hSensorPoint) {
      state.hSensorZero.push(hSensorPoint)
    },
    RESET(state) {
      state.pSensorLength = 0;
      state.tSensorLength = 0;
      state.hSensorLength = 0;
      state.pSensorZero = [null];
      state.tSensorZero = [null];
      state.hSensorZero = [null];
    }
  },
  actions: {
    setDevice({ commit }, device) {
      commit("SET_DEVICE", device)
    },
    setDeviceName({ commit }, deviceName) {
      commit("SET_DEVICENAME", deviceName)
    },
    setPSensor({ commit }, pSensor) {
      commit('SET_PSENSOR', pSensor)
    },
    setTSensor({ commit }, tSensor) {
      commit('SET_TSENSOR', tSensor)
    },
    setHSensor({ commit }, hSensor) {
      commit('SET_HSENSOR', hSensor)
    },
    setPSensorZero({ commit }, pSensorPoint) {
      commit('APPEND_PSENSORZERO', pSensorPoint)
    },
    setTSensorZero({ commit }, tSensorPoint) {
      commit('APPEND_TSENSORZERO', tSensorPoint)
    },
    setHSensorZero({ commit }, hSensorPoint) {
      commit('APPEND_HSENSORZERO', hSensorPoint)
    },
    reset({ commit }) {
      commit('RESET')
    }
  },
  getters: {
    selectedDevice(state) {
      return state.selectedDevice
    },
    selectedDeviceName(state) {
      return state.selectedDeviceName
    },
    url(state) {
      return state.url
    },
    pSensorLength(state) {
      return state.pSensorLength
    },
    tSensorLength(state) {
      return state.tSensorLength
    },
    hSensorLength(state) {
      return state.hSensorLength
    },
    pSensorZero(state) {
      return state.pSensorZero
    },
    tSensorZero(state) {
      return state.tSensorZero
    },
    hSensorZero(state) {
      return state.hSensorZero
    },
  },
  plugins: [
    createPersistedState()
  ]
})