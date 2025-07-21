// TODO: REMOVE THIS AS IT GETS CONVERTED INTO THE APPROPRATE PACKAGES.

package models

import (
    "time"
)

// Base model with common fields
type BaseModel struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}



// InventoryItem represents an item/part in the inventory.
type InventoryItem struct {
    BaseModel
    Name        string    `json:"name" db:"name"`
    Description string    `json:"description" db:"description"`
    SKU         string    `json:"sku" db:"sku"`
    Quantity    int       `json:"quantity" db:"quantity"`
    UnitPrice   float64   `json:"unit_price" db:"unit_price"`
    Supplier    string    `json:"supplier" db:"supplier"`
    Location    string    `json:"location" db:"location"`
	CostPrice     float64 `json:"cost_price" db:"cost_price"`
	SellPrice     float64 `json:"sell_price" db:"sell_price" binding:"required"`
	Category      string  `json:"category" binding:"required"`
}

// Customer
type Customer struct {
    BaseModel
    FirstName   string    `json:"first_name" db:"first_name"`
    LastName    string    `json:"last_name" db:"last_name"`
    Email       string    `json:"email" db:"email"`
    Phone       string    `json:"phone" db:"phone"`
    Address     string    `json:"address" db:"address"`
    City        string    `json:"city" db:"city"`
    PostalCode  string    `json:"postal_code" db:"postal_code"`

    // Relationships
	CustomerVehicles []CustomerVehicle `json:"customer_vehicles,omitempty"`
	Jobs     []Job     `json:"jobs,omitempty"`
}

// Vehicle represents vehicle data. Referenced by Customer Vehicle.
type Vehicle struct {
    BaseModel
    Make            string    `json:"make" db:"make"`
    Model           string    `json:"model" db:"model"`
    Year            int       `json:"year" db:"year"`
    ChassisNumber   string    `json:"chassis_number" db:"chassis_number"`
}

// A conjoining table betwixt a customer and their vehicles.
type CustomerVehicle struct {
    BaseModel
    VehicleID    int       `json:"vehicle_id" db:"vehicle_id"`
    CustomerID   int       `json:"customer_id" db:"customer_id"`
    VIN          string    `json:"vin" db:"vin"`
    License      string    `json:"license" db:"license"`
    Colour       string    `json:"colour" db:"colour"`
    Mileage      int       `json:"mileage" db:"mileage"`
    Notes        string    `json:"notes" db: "notes"`
    
    // Relationships
    Vehicle  Vehicle  `json:"vehicle`
	Customer Customer `json:"customer,omitempty"`
	Jobs     []Job    `json:"jobs,omitempty"`
}

// Job represents a work order/job
type Job struct {
    BaseModel
    CustomerID          int        `json:"customer_id" db:"customer_id"`
    CustomerVehicleID   int        `json:"vehicle_id" db:"vehicle_id"`
    Title               string     `json:"title" db:"title"`
    Description         string     `json:"description" db:"description"`
    Status              string     `json:"status" db:"status"`
    Priority            string     `json:"priority" db:"priority"`
    EstimatedCost       float64    `json:"estimated_cost" db:"estimated_cost"`
    ActualCost          float64    `json:"actual_cost" db:"actual_cost"`
    LaborHours          float64    `json:"labor_hours" db:"labor_hours"`
    StartDate           *time.Time `json:"start_date" db:"start_date"`
    EndDate             *time.Time `json:"end_date" db:"end_date"`

    // Relationships
	Customer        Customer        `json:"customer,omitempty"`
	CustomerVehicle CustomerVehicle `json:"customer_vehicle,omitempty"`
	AssignedTo      *User           `json:"assigned_to,omitempty"`
	JobItems        []JobItem       `json:"job_items,omitempty"`
	Payments        []Payment       `json:"payments,omitempty"`
}

// JobItem represents parts used in a job
type JobItem struct {
    BaseModel
    JobID         int     `json:"job_id" db:"job_id"`
    InventoryID   int     `json:"inventory_id" db:"inventory_id"`
    Quantity      int     `json:"quantity" db:"quantity"`
    UnitPrice     float64 `json:"unit_price" db:"unit_price"`
    TotalPrice    float64 `json:"total_price" db:"total_price"`

    // Relationships
	Job       Job            `json:"job,omitempty"`
	Inventory *InventoryItem `json:"inventory,omitempty"`
}

// Payment represents a payment record
type Payment struct {
    BaseModel
    JobID         int       `json:"job_id" db:"job_id"`
    Amount        float64   `json:"amount" db:"amount"`
    PaymentMethod string    `json:"payment_method" db:"payment_method"`
    PaymentDate   time.Time `json:"payment_date" db:"payment_date"`
    Status        string    `json:"status" db:"status"`
    Reference     string    `json:"reference" db:"reference"`
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