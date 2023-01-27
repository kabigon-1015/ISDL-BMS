import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueHead from 'vue-head'

const app = createApp(App)
app.use(VueHead)
app.use(router)
app.mount('#app')
