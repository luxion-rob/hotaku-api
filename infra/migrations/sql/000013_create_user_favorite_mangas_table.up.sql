CREATE TABLE user_favorite_mangas (
    favorite_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    user_id BIGINT NOT NULL,
    manga_id INT NOT NULL,
    added_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (favorite_id),
    INDEX idx_user_fav_user_id (user_id),
    INDEX idx_user_fav_manga_id (manga_id),
    CONSTRAINT fk_user_favorite_mangas_user FOREIGN KEY (user_id) REFERENCES users(user_id),
    CONSTRAINT fk_user_favorite_mangas_manga FOREIGN KEY (manga_id) REFERENCES mangas(manga_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;