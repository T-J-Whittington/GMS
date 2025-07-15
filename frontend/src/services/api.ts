import axios, { AxiosInstance, AxiosResponse } from 'axios'
import type {
  LoginRequest,
  RegisterRequest,
  LoginResponse,
  User,
  InventoryItem,
  CreateInventoryItem,
  Job,
  CreateJob,
  Payment,
  CreatePayment
} from '@/types'

class ApiService {
  private api: AxiosInstance

  constructor() {
    this.api = axios.create({
      baseURL: import.meta.env.VITE_API_URL || '/api/v1',
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json'
      }
    })

    // Request interceptor to add auth token
    this.api.interceptors.request.use(
      (config) => {
        const token = localStorage.getItem('token')
        if (token) {
          config.headers.Authorization = `Bearer ${token}`
        }
        return config
      },
      (error) => Promise.reject(error)
    )

    // Response interceptor for error handling
    this.api.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          localStorage.removeItem('token')
          window.location.href = '/login'
        }
        return Promise.reject(error)
      }
    )
  }

  // Helper method to handle API responses
  private handleResponse<T>(response: AxiosResponse<T>): T {
    return response.data
  }

  // Auth API
  async login(credentials: LoginRequest): Promise<LoginResponse> {
    const response = await this.api.post<LoginResponse>('/auth/login', credentials)
    return this.handleResponse(response)
  }

  async register(userData: RegisterRequest): Promise<LoginResponse> {
    const response = await this.api.post<LoginResponse>('/auth/register', userData)
    return this.handleResponse(response)
  }

  async getProfile(): Promise<User> {
    const response = await this.api.get<User>('/users/profile')
    return this.handleResponse(response)
  }

  // Inventory API
  async getInventory(): Promise<InventoryItem[]> {
    const response = await this.api.get<InventoryItem[]>('/inventory')
    return this.handleResponse(response)
  }

  async getInventoryItem(id: number): Promise<InventoryItem> {
    const response = await this.api.get<InventoryItem>(`/inventory/${id}`)
    return this.handleResponse(response)
  }

  async createInventoryItem(item: CreateInventoryItem): Promise<InventoryItem> {
    const response = await this.api.post<InventoryItem>('/inventory', item)
    return this.handleResponse(response)
  }

  async updateInventoryItem(id: number, item: Partial<InventoryItem>): Promise<InventoryItem> {
    const response = await this.api.put<InventoryItem>(`/inventory/${id}`, item)
    return this.handleResponse(response)
  }

  async deleteInventoryItem(id: number): Promise<void> {
    await this.api.delete(`/inventory/${id}`)
  }

  // Jobs API
  async getJobs(): Promise<Job[]> {
    const response = await this.api.get<Job[]>('/jobs')
    return this.handleResponse(response)
  }

  async getJob(id: number): Promise<Job> {
    const response = await this.api.get<Job>(`/jobs/${id}`)
    return this.handleResponse(response)
  }

  async createJob(job: CreateJob): Promise<Job> {
    const response = await this.api.post<Job>('/jobs', job)
    return this.handleResponse(response)
  }

  async updateJob(id: number, job: Partial<Job>): Promise<Job> {
    const response = await this.api.put<Job>(`/jobs/${id}`, job)
    return this.handleResponse(response)
  }

  async deleteJob(id: number): Promise<void> {
    await this.api.delete(`/jobs/${id}`)
  }

  // Payments API
  async getPayments(): Promise<Payment[]> {
    const response = await this.api.get<Payment[]>('/payments')
    return this.handleResponse(response)
  }

  async getPayment(id: number): Promise<Payment> {
    const response = await this.api.get<Payment>(`/payments/${id}`)
    return this.handleResponse(response)
  }

  async createPayment(payment: CreatePayment): Promise<Payment> {
    const response = await this.api.post<Payment>('/payments', payment)
    return this.handleResponse(response)
  }

  async updatePayment(id: number, payment: Partial<Payment>): Promise<Payment> {
    const response = await this.api.put<Payment>(`/payments/${id}`, payment)
    return this.handleResponse(response)
  }
}

// Create and export API service instance
const apiService = new ApiService()

// Export individual API modules for better organization
export const authApi = {
  login: apiService.login.bind(apiService),
  register: apiService.register.bind(apiService),
  getProfile: apiService.getProfile.bind(apiService)
}

export const inventoryApi = {
  getAll: apiService.getInventory.bind(apiService),
  getById: apiService.getInventoryItem.bind(apiService),
  create: apiService.createInventoryItem.bind(apiService),
  update: apiService.updateInventoryItem.bind(apiService),
  delete: apiService.deleteInventoryItem.bind(apiService)
}

export const jobsApi = {
  getAll: apiService.getJobs.bind(apiService),
  getById: apiService.getJob.bind(apiService),
  create: apiService.createJob.bind(apiService),
  update: apiService.updateJob.bind(apiService),
  delete: apiService.deleteJob.bind(apiService)
}

export const paymentsApi = {
  getAll: apiService.getPayments.bind(apiService),
  getById: apiService.getPayment.bind(apiService),
  create: apiService.createPayment.bind(apiService),
  update: apiService.updatePayment.bind(apiService)
}

export default apiService