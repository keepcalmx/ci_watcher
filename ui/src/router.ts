import { createRouter, createWebHistory } from 'vue-router'
import Overview from './views/Overview.vue'
import TestCase from './views/TestCase.vue'
import CaseLog from './views/CaseLog.vue'

const routes = [
  {
    path: '/',
    name: 'Overview',
    component: Overview
  },
  {
    path: '/cases',
    name: 'TestCase',
    component: TestCase
  },
  {
    path: '/logs',
    name: 'CaseLog',
    component: CaseLog,
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
