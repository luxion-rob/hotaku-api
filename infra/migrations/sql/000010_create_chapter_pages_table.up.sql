CREATE TABLE `chapter_pages` (
    page_id CHAR(36) NOT NULL,
    external_id CHAR(36) NOT NULL UNIQUE,
    chapter_id CHAR(36) NOT NULL,
    page_number INT UNSIGNED NOT NULL CHECK (page_number > 0),
    image_url VARCHAR(2048) NOT NULL,
    PRIMARY KEY (page_id),
    UNIQUE KEY uq_chapter_pages_chapter_id_page_number (chapter_id, page_number),
    CONSTRAINT fk_chapter_pages_chapter FOREIGN KEY (chapter_id) REFERENCES manga_chapters(chapter_id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
