CREATE TABLE mangas_groups (
    manga_id INT NOT NULL,
    group_id INT NOT NULL,
    PRIMARY KEY (manga_id, group_id),
    INDEX idx_mangas_groups_manga_id (manga_id),
    INDEX idx_mangas_groups_group_id (group_id),
    CONSTRAINT fk_mangas_groups_mangas FOREIGN KEY (manga_id) REFERENCES mangas(manga_id),
    CONSTRAINT fk_mangas_groups_groups FOREIGN KEY (group_id) REFERENCES `groups`(group_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;