CREATE TABLE
    `user_favorite_mangas` (
        favorite_id INT NOT NULL AUTO_INCREMENT,
        external_id CHAR(36),
        user_id BIGINT UNSIGNED,
        manga_id INT,
        added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        primary KEY (favorite_id),
        foreign KEY (user_id) REFERENCES `users` (id) ON DELETE CASCADE,
        foreign KEY (manga_id) REFERENCES `mangas` (manga_id) ON DELETE CASCADE
    );