CREATE TABLE `mangas` (
    manga_id CHAR(36) NOT NULL,
    external_id CHAR(36) NOT NULL UNIQUE,
    status_id INT UNSIGNED NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (manga_id),
        CONSTRAINT fk_mangas_manga_statuses_status_id
        FOREIGN KEY (status_id) REFERENCES manga_statuses(status_id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    INDEX idx_mangas_status (status_id),
    INDEX idx_mangas_created_at (created_at)
    FULLTEXT KEY ft_mangas_title_description (title)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
