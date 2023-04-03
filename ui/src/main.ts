import { createApp } from 'vue'
import App from './App.vue'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import VueApexCharts from "vue3-apexcharts";

const app = createApp(App)

app.use(ElementPlus)
app.use(VueApexCharts)
app.mount('#app')
