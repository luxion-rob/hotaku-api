CREATE TABLE mangas_authors (
    manga_id VARCHAR(36) NOT NULL,
    author_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (manga_id, author_id),
    CONSTRAINT fk_mangas_authors_mangas FOREIGN KEY (manga_id) REFERENCES mangas(manga_id) ON DELETE CASCADE,
    CONSTRAINT fk_mangas_authors_authors FOREIGN KEY (author_id) REFERENCES authors(author_id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;