import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueHead from 'vue-head'

Vue.use(VueHead)
createApp(App).use(router).mount('#app')

