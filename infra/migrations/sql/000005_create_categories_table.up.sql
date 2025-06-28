CREATE TABLE categories (
    category_id VARCHAR(36) NOT NULL UNIQUE,
    external_id VARCHAR(36) NOT NULL UNIQUE,
    category_name VARCHAR(40) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (category_id),
    UNIQUE KEY uq_categories_name (category_name)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;