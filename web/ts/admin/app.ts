import { createApp } from 'vue/dist/vue.esm-bundler'
import App from './App.vue'
import router from './router'
import Pusher from 'pusher-js'
import globalValues from '@/plugins/globalValues'

// @ts-ignore
window.Pusher = Pusher

const app = createApp(App)

app.use(router).use(globalValues).mount('#admin-main')
