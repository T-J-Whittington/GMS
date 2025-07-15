import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/dashboard'
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/LoginView.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('@/views/RegisterView.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: () => import('@/views/DashboardView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/inventory',
      name: 'Inventory',
      component: () => import('@/views/InventoryView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/jobs',
      name: 'Jobs',
      component: () => import('@/views/JobsView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/customers',
      name: 'Customers',
      component: () => import('@/views/CustomersView.vue'),
      meta: { requiresAuth: true }
    },
    {
      path: '/payments',
      name: 'Payments',
      component: () => import('@/views/PaymentsView.vue'),
      meta: { requiresAuth: true }
    }
  ]
})

// Navigation guards
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
  } else if (to.meta.requiresGuest && authStore.isAuthenticated) {
    next('/dashboard')
  } else {
    next()
  }
})

export default router