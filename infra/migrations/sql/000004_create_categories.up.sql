CREATE TABLE `categories` (
    category_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    external_id CHAR(36) NOT NULL UNIQUE,
    category_name VARCHAR(255) NOT NULL
)