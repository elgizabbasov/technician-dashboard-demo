import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('../views/Dashboard.vue')
  },
  {
    path: '/form',
    name: 'form',
    component: () => import('../views/TechnicianForm.vue')
  },
  {
    path: '/setup',
    name: 'setup',
    component: () => import('../views/Setup.vue')
  },
  {
    path: '/',
    name: 'selectdevice',
    component: () => import('../views/SelectDevice.vue')
  }

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router;
