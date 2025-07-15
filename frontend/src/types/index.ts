// User types
export interface User {
  id: number
  email: string
  first_name: string
  last_name: string
  role: string
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  email: string
  password: string
  first_name: string
  last_name: string
}

export interface LoginResponse {
  token: string
  user: User
}

// Inventory types
export interface InventoryItem {
  id: number
  name: string
  description: string
  sku: string
  quantity: number
  min_stock: number
  unit_price: number
  supplier: string
  location: string
  created_at: string
  updated_at: string
}

export interface CreateInventoryItem {
  name: string
  description: string
  sku: string
  quantity: number
  min_stock: number
  unit_price: number
  supplier: string
  location: string
}

// Customer types
export interface Customer {
  id: number
  first_name: string
  last_name: string
  email: string
  phone: string
  address: string
  city: string
  postal_code: string
  created_at: string
  updated_at: string
}

export interface CreateCustomer {
  first_name: string
  last_name: string
  email: string
  phone: string
  address: string
  city: string
  postal_code: string
}

// Vehicle types
export interface Vehicle {
  id: number
  customer_id: number
  make: string
  model: string
  year: number
  vin: string
  license_plate: string
  color: string
  mileage: number
  created_at: string
  updated_at: string
}

// Job types
export interface Job {
  id: number
  customer_id: number
  vehicle_id: number
  title: string
  description: string
  status: string
  priority: string
  estimated_cost: number
  actual_cost: number
  labor_hours: number
  start_date: string | null
  end_date: string | null
  created_at: string
  updated_at: string
}

export interface CreateJob {
  customer_id: number
  vehicle_id: number
  title: string
  description: string
  status: string
  priority: string
  estimated_cost: number
  labor_hours: number
}

// Payment types
export interface Payment {
  id: number
  job_id: number
  amount: number
  payment_method: string
  payment_date: string
  status: string
  reference: string
  created_at: string
  updated_at: string
}

export interface CreatePayment {
  job_id: number
  amount: number
  payment_method: string
  reference: string
}

// API Response types
export interface ApiResponse<T> {
  data: T
  message?: string
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  limit: number
  pages: number
}