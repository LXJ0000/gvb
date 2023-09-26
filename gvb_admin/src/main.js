
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import Antd from 'ant-design-vue';
import App from './App.vue'
// import 'ant-design-vue/dist/reset.css';
import router from './router'
import './assets/css/public.css'
import './assets/css/iconfont.css'
import './assets/css/antdv.css'
import './assets/css/theme.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Antd)
app.mount('#app')



