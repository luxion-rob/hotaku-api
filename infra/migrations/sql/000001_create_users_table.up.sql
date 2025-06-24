CREATE TABLE `users` (
        id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        created_at DATETIME (3) NULL,
        updated_at DATETIME (3) NULL,
        deleted_at DATETIME (3) NULL,
        email VARCHAR(191) NOT NULL UNIQUE,
        password VARCHAR(255) NOT NULL,
        name VARCHAR(255) NOT NULL,
        primary KEY (id),
        index idx_users_deleted_at (deleted_at),
        index idx_users_email (email)
    );