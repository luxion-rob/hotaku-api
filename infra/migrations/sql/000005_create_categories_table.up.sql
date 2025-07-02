CREATE TABLE categories (
    category_id CHAR(36) NOT NULL,
    external_id CHAR(36) NOT NULL,
    category_name VARCHAR(40) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (category_id),
    UNIQUE KEY uq_categories_category_name (category_name)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;