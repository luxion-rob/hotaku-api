CREATE TABLE mangas (
    manga_id VARCHAR(36) NOT NULL UNIQUE,
    external_id VARCHAR(36) NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status ENUM('ongoing', 'completed', 'hiatus', 'cancelled') NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (manga_id),
    INDEX idx_status (status),
    INDEX idx_manga_title (title)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;