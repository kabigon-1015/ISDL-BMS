import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueHead from 'vue-head'
import store from './store'
// import { GlobalStoreKey, globalStore } from './store/global.js'
// app.config.globalProperties.$username = "岡颯人";
// app.provide(GlobalStoreKey, globalStore())
const app = createApp(App).use(store)
app.use(VueHead)
app.use(router)
app.use(store)
app.mount('#app')
