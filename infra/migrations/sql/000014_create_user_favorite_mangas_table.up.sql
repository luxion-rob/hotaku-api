CREATE TABLE `user_favorite_mangas` (
    favorite_id CHAR(36) NOT NULL,
    external_id CHAR(36) NOT NULL UNIQUE,
    user_id CHAR(36) NOT NULL,
    manga_id CHAR(36) NOT NULL,
    added_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_user_favorite_mangas_added_at (added_at),
    INDEX idx_user_favorite_mangas_user_id_added_at (user_id, added_at DESC),
    PRIMARY KEY (favorite_id),
    UNIQUE KEY uq_user_favorite_mangas_user_id_manga_id (user_id, manga_id),
    CONSTRAINT fk_user_favorite_mangas_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_user_favorite_mangas_manga FOREIGN KEY (manga_id) REFERENCES mangas(manga_id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;