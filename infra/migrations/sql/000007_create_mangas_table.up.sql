CREATE TABLE mangas (
    manga_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    title VARCHAR(255),
    description TEXT,
    status ENUM('ongoing', 'completed', 'hiatus', 'cancelled') NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_flag BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP,
    PRIMARY KEY (manga_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;