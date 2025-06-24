CREATE TABLE `manga_chapters` (
        chapter_id INT NOT NULL AUTO_INCREMENT,
        external_id CHAR(36),
        manga_id INT,
        chapter_number INT NOT NULL,
        title VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        primary KEY (chapter_id),
        foreign KEY (manga_id) REFERENCES `mangas` (manga_id) ON DELETE SET NULL
    );