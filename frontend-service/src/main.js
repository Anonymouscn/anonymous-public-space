import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import MetaInfo from 'vue-meta-info';
import './assets/common/css/global.css'
import './assets/common/css/font.css'
import 'virtual:svg-icons-register'
import SvgIcon from './components/SvgIcon.vue'
import message from "./components/message/index.js";

const app = createApp(App).use(router).use(MetaInfo).component('svg-icon', SvgIcon)
app.config.globalProperties.$message = message
app.mount('#app')