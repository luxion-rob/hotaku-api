CREATE TABLE users (
    user_id varchar(36) NOT NULL UNIQUE,
    role_id varchar(36) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_flag BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP,
    PRIMARY KEY (user_id),
    INDEX idx_users_role_id (role_id),
    INDEX idx_users_deleted_at (deleted_at),
    CONSTRAINT fk_users_roles FOREIGN KEY (role_id) REFERENCES roles(role_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;