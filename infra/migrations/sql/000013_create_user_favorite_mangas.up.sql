CREATE TABLE `user_favorite_mangas` (
    favorite_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    user_id BIGINT UNSIGNED NOT NULL,
    manga_id INT NOT NULL,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    primary KEY (favorite_id),
    UNIQUE KEY uq_external_id (external_id),
    UNIQUE KEY uq_user_manga (user_id, manga_id),
    foreign KEY (user_id) REFERENCES `users` (id) ON DELETE CASCADE,
    foreign KEY (manga_id) REFERENCES `mangas` (manga_id) ON DELETE CASCADE
);