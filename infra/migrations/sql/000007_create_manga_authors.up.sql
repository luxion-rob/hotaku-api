CREATE TABLE `manga_authors` (
        manga_id INT NOT NULL,
        author_id INT NOT NULL,
        primary KEY (manga_id, author_id),
        foreign KEY (manga_id) REFERENCES `mangas` (manga_id) ON DELETE CASCADE,
        foreign KEY (author_id) REFERENCES `authors` (author_id) ON DELETE CASCADE
    );