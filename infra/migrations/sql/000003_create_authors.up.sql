CREATE TABLE `authors` (
    author_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    author_name VARCHAR(255),
    author_bio TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary KEY (author_id)
);
CREATE INDEX idx_authors_name ON `authors` (author_name);