CREATE TABLE manga_chapters (
    chapter_id CHAR(36) NOT NULL,
    external_id CHAR(36) NOT NULL,
    manga_id CHAR(36) NOT NULL,
    chapter_number DECIMAL(10, 2) NOT NULL,
    title VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (chapter_id),
    UNIQUE KEY uq_manga_chapters_manga_id_chapter_number (manga_id, chapter_number),
    CONSTRAINT fk_manga_chapters_mangas FOREIGN KEY (manga_id) REFERENCES mangas(manga_id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;