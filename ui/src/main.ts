import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)

// 路由管理
import router from './router'
app.use(router)

// ElementPlus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
app.use(ElementPlus)

import * as ElementPlusIconsVue from '@element-plus/icons-vue'
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

// VueApexCharts
import VueApexCharts from "vue3-apexcharts";
app.use(VueApexCharts)

app.mount('#app')
