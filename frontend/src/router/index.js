import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Dashboard from '../views/Dashboard.vue'
import PaperList from '../views/PaperList.vue'
import AddPaper from '../views/AddPaper.vue'
import PaperDetail from '../views/PaperDetail.vue'
import AddQuestion from '../views/AddQuestion.vue'
import QuestionList from '../views/QuestionList.vue'
import Favorites from '../views/Favorites.vue'
import Mistakes from '../views/Mistakes.vue'
import LearningHistory from '../views/LearningHistory.vue'
import Notes from '../views/Notes.vue'
import Discussion from '../views/Discussion.vue'
import AdminDashboard from '../views/admin/AdminDashboard.vue'
import UserManage from '../views/admin/UserManage.vue'
import Help from '../views/Help.vue'
import About from '../views/About.vue'
import Notifications from '../views/Notifications.vue'
import NotFound from '../views/NotFound.vue'
import Membership from '../views/Membership.vue'

const routes = [
  { path: '/login', name: 'Login', component: Login },
  { path: '/register', name: 'Register', component: Register },
  { path: '/', name: 'Dashboard', component: Dashboard },
  { path: '/papers', name: 'PaperList', component: PaperList },
  { path: '/papers/add', name: 'AddPaper', component: AddPaper },
  { path: '/papers/:id', name: 'PaperDetail', component: PaperDetail },
  { path: '/papers/:id/questions/add', name: 'AddQuestion', component: AddQuestion },
  { path: '/questions', name: 'QuestionList', component: QuestionList },
  { path: '/favorites', name: 'Favorites', component: Favorites },
  { path: '/mistakes', name: 'Mistakes', component: Mistakes },
  { path: '/history', name: 'LearningHistory', component: LearningHistory },
  { path: '/notes', name: 'Notes', component: Notes },
  { path: '/discussion', name: 'Discussion', component: Discussion },
  { path: '/admin', name: 'AdminDashboard', component: AdminDashboard },
  { path: '/admin/users', name: 'UserManage', component: UserManage },
  { path: '/help', name: 'Help', component: Help },
  { path: '/about', name: 'About', component: About },
  { path: '/notifications', name: 'Notifications', component: Notifications },
  { path: '/membership', name: 'Membership', component: Membership },
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫：检查登录状态
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  // 如果没有token且不是登录或注册页面，则跳转到登录页
  if (!token && to.path !== '/login' && to.path !== '/register') {
    next('/login')
  } else {
    next()
  }
})

export default router 