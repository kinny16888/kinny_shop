import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from '../router'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import "core-js/stable";
import "regenerator-runtime/runtime";
import * as am4core from "@amcharts/amcharts4/core";
import * as am4charts from "@amcharts/amcharts4/charts";
import am4themes_animated from "@amcharts/amcharts4/themes/animated";

const app = createApp(App)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
window.router=router
app.use(router)
app.use(ElementPlus)
app.config.globalProperties.am4core = am4core
app.config.globalProperties.am4charts = am4charts
app.config.globalProperties.am4themes_animated = am4themes_animated
app.mount('#app')
