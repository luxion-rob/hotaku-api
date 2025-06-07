-- Test database initialization script
-- This script runs when the MySQL container starts up

-- Ensure we're using the test database
USE hotaku_test_db;

-- Create test users if they don't exist
CREATE USER IF NOT EXISTS 'testuser'@'%' IDENTIFIED BY 'testpassword';
GRANT ALL PRIVILEGES ON hotaku_test_db.* TO 'testuser'@'%';

-- Sample table structure (adjust based on your actual schema)
-- You can add your actual table definitions here

-- Example: Users table for testing
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Example: Insert test data
INSERT IGNORE INTO users (email, password_hash) VALUES 
('test@example.com', 'hashed_password_here'),
('admin@example.com', 'admin_hashed_password_here');

-- Flush privileges
FLUSH PRIVILEGES; 