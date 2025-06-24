CREATE TABLE `user_manga_histories` (
    history_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    user_id BIGINT UNSIGNED,
    manga_id INT,
    read_chapter_ids TEXT,
    read_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    primary KEY (history_id),
    foreign KEY (user_id) REFERENCES `users` (id) ON DELETE CASCADE,
    foreign KEY (manga_id) REFERENCES `mangas` (manga_id) ON DELETE CASCADE
);