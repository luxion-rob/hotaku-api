CREATE TABLE chapter_pages (
    page_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    chapter_id INT NOT NULL,
    page_number INT,
    image_url VARCHAR(500) NOT NULL,
    PRIMARY KEY (page_id),
    INDEX idx_chapter_pages_chapter_id (chapter_id),
    CONSTRAINT fk_chapter_pages_chapter FOREIGN KEY (chapter_id) REFERENCES manga_chapters(chapter_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;