CREATE TABLE chapter_pages (
    page_id VARCHAR(36) NOT NULL UNIQUE,
    external_id VARCHAR(36) NOT NULL UNIQUE,
    chapter_id VARCHAR(36) NOT NULL,
    page_number INT,
    image_url VARCHAR(500) NOT NULL,
    PRIMARY KEY (page_id),
    INDEX idx_chapter_pages_chapter_id (chapter_id),
    UNIQUE KEY uq_chapter_pages_number (chapter_id, page_number),
    CONSTRAINT fk_chapter_pages_chapter FOREIGN KEY (chapter_id) REFERENCES manga_chapters(chapter_id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;