import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Dashboard from '../views/Dashboard.vue'
import PaperList from '../views/PaperList.vue'

const routes = [
  { path: '/login', name: 'Login', component: Login },
  { path: '/register', name: 'Register', component: Register },
  { path: '/', name: 'Dashboard', component: Dashboard },
  { path: '/papers', name: 'PaperList', component: PaperList },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 