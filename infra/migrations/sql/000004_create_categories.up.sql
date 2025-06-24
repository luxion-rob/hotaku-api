CREATE TABLE `categories` (
    category_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36),
    category_name VARCHAR(255),
    primary KEY (category_id)
);
