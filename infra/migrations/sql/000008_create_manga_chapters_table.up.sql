CREATE TABLE manga_chapters (
    chapter_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    manga_id INT NOT NULL,
    chapter_number INT NOT NULL,
    title VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_flag BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP,
    PRIMARY KEY (chapter_id),
    INDEX idx_manga_chapters_manga_id (manga_id),
    CONSTRAINT fk_manga_chapters_mangas FOREIGN KEY (manga_id) REFERENCES mangas(manga_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;