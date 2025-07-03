CREATE TABLE `mangas_categories` (
    manga_id CHAR(36) NOT NULL,
    category_id CHAR(36) NOT NULL,
    PRIMARY KEY (manga_id, category_id),
    INDEX idx_mangas_categories_category_id (category_id),
    CONSTRAINT fk_mangas_categories_mangas FOREIGN KEY (manga_id) REFERENCES mangas(manga_id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_mangas_categories_categories FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
