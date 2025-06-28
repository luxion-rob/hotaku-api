CREATE TABLE mangas_authors (
    manga_id INT NOT NULL,
    author_id INT NOT NULL,
    PRIMARY KEY (manga_id, author_id),
    INDEX idx_mangas_authors_manga_id (manga_id),
    INDEX idx_mangas_authors_author_id (author_id),
    CONSTRAINT fk_mangas_authors_mangas FOREIGN KEY (manga_id) REFERENCES mangas(manga_id),
    CONSTRAINT fk_mangas_authors_authors FOREIGN KEY (author_id) REFERENCES authors(author_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;