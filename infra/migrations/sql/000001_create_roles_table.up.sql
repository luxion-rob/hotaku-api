CREATE TABLE `roles` (
    role_id CHAR(36) NOT NULL,
    role_name VARCHAR(100) NOT NULL UNIQUE,
    PRIMARY KEY (role_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;