CREATE TABLE user_read_chapters (
    user_id CHAR(36) NOT NULL,
    chapter_id CHAR(36) NOT NULL,
    read_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, chapter_id),
    INDEX idx_user_read_chapters_chapter_id (chapter_id),
    CONSTRAINT fk_user_read_chapters_users
        FOREIGN KEY (user_id) REFERENCES users(user_id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_user_read_chapters_chapters
        FOREIGN KEY (chapter_id) REFERENCES manga_chapters(chapter_id)
        ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;