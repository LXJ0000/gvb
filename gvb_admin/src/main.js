import {createApp} from 'vue'
import Antd from 'ant-design-vue';
import {createPinia} from 'pinia'
import App from './App.vue'
import 'ant-design-vue/dist/reset.css';
import '@/assets/css/theme.css'
import '@/assets/css/public.css'
import '@/assets/css/iconfont.css'

import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Antd)
app.mount('#app')
