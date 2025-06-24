CREATE TABLE `manga_categories` (
    manga_id INT NOT NULL,
    category_id INT NOT NULL,
    primary KEY (manga_id, category_id),
    foreign KEY (manga_id) REFERENCES `mangas` (manga_id) ON DELETE CASCADE,
    foreign KEY (category_id) REFERENCES `categories` (category_id) ON DELETE CASCADE
);
CREATE INDEX idx_category_manga ON manga_categories (category_id, manga_id);