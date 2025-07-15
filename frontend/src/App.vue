<template>
  <div id="app">
    <nav class="navbar" v-if="authStore.isAuthenticated">
      <div class="navbar-brand">
        <router-link to="/" class="brand-link">
          🔧 Garage Manager
        </router-link>
      </div>
      
      <div class="navbar-menu">
        <router-link to="/dashboard" class="nav-link">Dashboard</router-link>
        <router-link to="/inventory" class="nav-link">Inventory</router-link>
        <router-link to="/jobs" class="nav-link">Jobs</router-link>
        <router-link to="/customers" class="nav-link">Customers</router-link>
        <router-link to="/payments" class="nav-link">Payments</router-link>
      </div>

      <div class="navbar-user">
        <span class="user-name">{{ userFullName }}</span>
        <button @click="logout" class="logout-btn">Logout</button>
      </div>
    </nav>

    <main class="main-content" :class="{ 'with-navbar': authStore.isAuthenticated }">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const userFullName = computed(() => {
  const user = authStore.user
  return user ? `${user.first_name} ${user.last_name}` : ''
})

const logout = async () => {
  await authStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.navbar {
  background: #2c3e50;
  color: white;
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.navbar-brand .brand-link {
  font-size: 1.5rem;
  font-weight: bold;
  color: white;
  text-decoration: none;
}

.navbar-menu {
  display: flex;
  gap: 2rem;
}

.nav-link {
  color: white;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.nav-link:hover,
.nav-link.router-link-active {
  background-color: #34495e;
}

.navbar-user {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-name {
  font-weight: 500;
}

.logout-btn {
  background: #e74c3c;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.logout-btn:hover {
  background: #c0392b;
}

.main-content {
  min-height: 100vh;
  background: #f8f9fa;
}

.main-content.with-navbar {
  min-height: calc(100vh - 70px);
}

#app {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}
</style>