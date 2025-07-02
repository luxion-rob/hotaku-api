CREATE TABLE users (
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
DELIMITER $$ -- Trigger trước khi insert
CREATE TRIGGER trg_users_bi_set_deleted_flag BEFORE
INSERT ON `users` FOR EACH ROW BEGIN IF NEW.deleted_at IS NOT NULL THEN
SET NEW.deleted_flag = TRUE;
ELSE
SET NEW.deleted_flag = FALSE;
END IF;
END $$ -- Trigger trước khi update
CREATE TRIGGER trg_users_bu_set_deleted_flag BEFORE
UPDATE ON `users` FOR EACH ROW BEGIN IF NEW.deleted_at IS NOT NULL THEN
SET NEW.deleted_flag = TRUE;
ELSE
SET NEW.deleted_flag = FALSE;
END IF;
END $$ DELIMITER;