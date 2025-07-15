package models

import (
    "time"
)

// User represents a system user
type User struct {
    ID        int       `json:"id" db:"id"`
    Email     string    `json:"email" db:"email"`
    Password  string    `json:"-" db:"password"`
    FirstName string    `json:"first_name" db:"first_name"`
    LastName  string    `json:"last_name" db:"last_name"`
    Role      string    `json:"role" db:"role"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// InventoryItem represents a parts inventory item
type InventoryItem struct {
    ID          int       `json:"id" db:"id"`
    Name        string    `json:"name" db:"name"`
    Description string    `json:"description" db:"description"`
    SKU         string    `json:"sku" db:"sku"`
    Quantity    int       `json:"quantity" db:"quantity"`
    MinStock    int       `json:"min_stock" db:"min_stock"`
    UnitPrice   float64   `json:"unit_price" db:"unit_price"`
    Supplier    string    `json:"supplier" db:"supplier"`
    Location    string    `json:"location" db:"location"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Customer represents a customer
type Customer struct {
    ID          int       `json:"id" db:"id"`
    FirstName   string    `json:"first_name" db:"first_name"`
    LastName    string    `json:"last_name" db:"last_name"`
    Email       string    `json:"email" db:"email"`
    Phone       string    `json:"phone" db:"phone"`
    Address     string    `json:"address" db:"address"`
    City        string    `json:"city" db:"city"`
    PostalCode  string    `json:"postal_code" db:"postal_code"`
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Vehicle represents a customer's vehicle
type Vehicle struct {
    ID           int       `json:"id" db:"id"`
    CustomerID   int       `json:"customer_id" db:"customer_id"`
    Make         string    `json:"make" db:"make"`
    Model        string    `json:"model" db:"model"`
    Year         int       `json:"year" db:"year"`
    VIN          string    `json:"vin" db:"vin"`
    LicensePlate string    `json:"license_plate" db:"license_plate"`
    Color        string    `json:"color" db:"color"`
    Mileage      int       `json:"mileage" db:"mileage"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
    UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Job represents a work order/job
type Job struct {
    ID          int       `json:"id" db:"id"`
    CustomerID  int       `json:"customer_id" db:"customer_id"`
    VehicleID   int       `json:"vehicle_id" db:"vehicle_id"`
    Title       string    `json:"title" db:"title"`
    Description string    `json:"description" db:"description"`
    Status      string    `json:"status" db:"status"`
    Priority    string    `json:"priority" db:"priority"`
    EstimatedCost float64 `json:"estimated_cost" db:"estimated_cost"`
    ActualCost    float64 `json:"actual_cost" db:"actual_cost"`
    LaborHours    float64 `json:"labor_hours" db:"labor_hours"`
    StartDate     *time.Time `json:"start_date" db:"start_date"`
    EndDate       *time.Time `json:"end_date" db:"end_date"`
    CreatedAt     time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

// JobItem represents parts used in a job
type JobItem struct {
    ID            int     `json:"id" db:"id"`
    JobID         int     `json:"job_id" db:"job_id"`
    InventoryID   int     `json:"inventory_id" db:"inventory_id"`
    Quantity      int     `json:"quantity" db:"quantity"`
    UnitPrice     float64 `json:"unit_price" db:"unit_price"`
    TotalPrice    float64 `json:"total_price" db:"total_price"`
}

// Payment represents a payment record
type Payment struct {
    ID            int       `json:"id" db:"id"`
    JobID         int       `json:"job_id" db:"job_id"`
    Amount        float64   `json:"amount" db:"amount"`
    PaymentMethod string    `json:"payment_method" db:"payment_method"`
    PaymentDate   time.Time `json:"payment_date" db:"payment_date"`
    Status        string    `json:"status" db:"status"`
    Reference     string    `json:"reference" db:"reference"`
    CreatedAt     time.Time `json:"created_at" db:"created_at"`
    UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// LoginRequest represents login request payload
type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// RegisterRequest represents registration request payload
type RegisterRequest struct {
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=6"`
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
}

// LoginResponse represents login response
type LoginResponse struct {
    Token string `json:"token"`
    User  User   `json:"user"`
}