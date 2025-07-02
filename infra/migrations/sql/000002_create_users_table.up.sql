CREATE TABLE `users` (
    user_id CHAR(36) NOT NULL,
    role_id CHAR(36) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_flag BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP,
    PRIMARY KEY (user_id),
    INDEX idx_users_role_id (role_id),
    INDEX idx_users_deleted_at (deleted_at),
    CONSTRAINT fk_users_roles FOREIGN KEY (role_id) REFERENCES roles(role_id) ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
