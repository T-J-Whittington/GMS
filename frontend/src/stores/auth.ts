import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi } from '@/services/api'
import type { User, LoginRequest, RegisterRequest } from '@/types'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(localStorage.getItem('token'))
  const loading = ref(false)

  const isAuthenticated = computed(() => !!token.value && !!user.value)

  const login = async (credentials: LoginRequest) => {
    loading.value = true
    try {
      const response = await authApi.login(credentials)
      token.value = response.token
      user.value = response.user
      localStorage.setItem('token', response.token)
      return response
    } catch (error) {
      throw error
    } finally {
      loading.value = false
    }
  }

  const register = async (userData: RegisterRequest) => {
    loading.value = true
    try {
      const response = await authApi.register(userData)
      token.value = response.token
      user.value = response.user
      localStorage.setItem('token', response.token)
      return response
    } catch (error) {
      throw error
    } finally {
      loading.value = false
    }
  }

  const logout = async () => {
    user.value = null
    token.value = null
    localStorage.removeItem('token')
  }

  const loadUser = async () => {
    if (!token.value) return

    try {
      const userData = await authApi.getProfile()
      user.value = userData
    } catch (error) {
      // Token might be invalid, clear it
      logout()
    }
  }

  // Initialize auth state
  if (token.value) {
    loadUser()
  }

  return {
    user,
    token,
    loading,
    isAuthenticated,
    login,
    register,
    logout,
    loadUser
  }
})