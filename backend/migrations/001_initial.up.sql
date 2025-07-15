-- Initial database schema for garage management system

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    role VARCHAR(50) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Customers table
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(20),
    address TEXT,
    city VARCHAR(100),
    postal_code VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Vehicles table
CREATE TABLE IF NOT EXISTS vehicles (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(id) ON DELETE CASCADE,
    make VARCHAR(100) NOT NULL,
    model VARCHAR(100) NOT NULL,
    year INTEGER NOT NULL,
    vin VARCHAR(50) UNIQUE,
    license_plate VARCHAR(20),
    color VARCHAR(50),
    mileage INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inventory items table
CREATE TABLE IF NOT EXISTS inventory_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    sku VARCHAR(100) UNIQUE NOT NULL,
    quantity INTEGER DEFAULT 0,
    min_stock INTEGER DEFAULT 0,
    unit_price DECIMAL(10,2) DEFAULT 0.00,
    supplier VARCHAR(255),
    location VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Jobs table
CREATE TABLE IF NOT EXISTS jobs (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(id) ON DELETE CASCADE,
    vehicle_id INTEGER REFERENCES vehicles(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'pending',
    priority VARCHAR(50) DEFAULT 'medium',
    estimated_cost DECIMAL(10,2) DEFAULT 0.00,
    actual_cost DECIMAL(10,2) DEFAULT 0.00,
    labor_hours DECIMAL(5,2) DEFAULT 0.00,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Job items table (parts used in jobs)
CREATE TABLE IF NOT EXISTS job_items (
    id SERIAL PRIMARY KEY,
    job_id INTEGER REFERENCES jobs(id) ON DELETE CASCADE,
    inventory_id INTEGER REFERENCES inventory_items(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    total_price DECIMAL(10,2) GENERATED ALWAYS AS (quantity * unit_price) STORED
);

-- Payments table
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    job_id INTEGER REFERENCES jobs(id) ON DELETE CASCADE,
    amount DECIMAL(10,2) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) DEFAULT 'completed',
    reference VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_customers_email ON customers(email);
CREATE INDEX IF NOT EXISTS idx_vehicles_customer_id ON vehicles(customer_id);
CREATE INDEX IF NOT EXISTS idx_inventory_sku ON inventory_items(sku);
CREATE INDEX IF NOT EXISTS idx_jobs_customer_id ON jobs(customer_id);
CREATE INDEX IF NOT EXISTS idx_jobs_vehicle_id ON jobs(vehicle_id);
CREATE INDEX IF NOT EXISTS idx_jobs_status ON jobs(status);
CREATE INDEX IF NOT EXISTS idx_job_items_job_id ON job_items(job_id);
CREATE INDEX IF NOT EXISTS idx_payments_job_id ON payments(job_id);

-- Create triggers for updated_at timestamps
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_customers_updated_at BEFORE UPDATE ON customers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_vehicles_updated_at BEFORE UPDATE ON vehicles FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_inventory_items_updated_at BEFORE UPDATE ON inventory_items FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_jobs_updated_at BEFORE UPDATE ON jobs FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_payments_updated_at BEFORE UPDATE ON payments FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();