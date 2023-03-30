import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import BookRental from '../views/BookRental.vue'
import BarcodeReader from '../views/BarcodeReader.vue'
import BookReturn from '../views/BookReturn.vue'
import UserLogin from '../views/UserLogin.vue'
import SignUp from '../views/SignUp.vue'
import ResetPassword from '../views/ResetPassword.vue'
import BookList from '../views/BookList.vue'
import BookResearch from '../views/BookResearch.vue'
import BookDetail from '../views/BookDetail.vue'
import AddBookTag from '../views/AddBookTag.vue'
import UserInfo from '../views/UserInfo.vue'
import ChangeUserInfo from '../views/ChangeUserInfo.vue'

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
  },
  {
    path: '/return',
    name: 'return',
    component: BookReturn,
  },
  {
    path: '/userlogin',
    name: 'UserLogin',
    component: UserLogin,
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: SignUp,
  },
  {
    path: '/resetpassword',
    name: 'ResetPassword',
    component: ResetPassword,
  },
  {
    path: '/booklist',
    name: 'booklist',
    component: BookList,
  },
  {
    path: '/rental/BookResearch',
    name: 'BookResearch',
    component: BookResearch,
  },
  {
    path: '/booklist/detail',
    name: 'detail',
    component: BookDetail,
  },
  {
    path: '/addbooktag',
    name: 'addbooktag',
    component: AddBookTag,
  },
  {
    path: '/userinfo',
    name: 'userinfo',
    component: UserInfo,
  },
  {
    path: '/changeuserinfo',
    name: 'changeuserinfo',
    component: ChangeUserInfo,
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
