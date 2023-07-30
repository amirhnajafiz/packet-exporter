import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import DocumentsView from '../views/Documents.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/website',
      name: 'home',
      component: HomeView
    },
    {
      path: '/website/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/website/docs',
      name: 'docs',
      component: DocumentsView
    }
  ]
})

export default router
