import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import BookRental from '../views/BookRental.vue'
import BarcodeReader from '../views/BarcodeReader.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/rental',
    name: 'rental',
    component: BookRental,
  },
  {
    path: '/rental/BarcodeReader',
    name: 'BarcodeReader',
    component: BarcodeReader,
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
