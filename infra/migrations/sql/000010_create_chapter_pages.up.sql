CREATE TABLE `chapter_pages` (
    page_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    chapter_id INT,
    image_url VARCHAR(500) NOT NULL,
    page_number INT,
    primary KEY (page_id),
    foreign KEY (chapter_id) REFERENCES `manga_chapters` (chapter_id) ON DELETE CASCADE
);