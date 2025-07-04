CREATE TABLE `mangas_groups` (
    manga_id CHAR(36) NOT NULL,
    group_id CHAR(36) NOT NULL,
    PRIMARY KEY (manga_id, group_id),
    INDEX idx_mangas_groups_group_id (group_id),
    CONSTRAINT fk_mangas_groups_mangas FOREIGN KEY (manga_id) REFERENCES mangas(manga_id) ON DELETE CASCADE,
    CONSTRAINT fk_mangas_groups_groups FOREIGN KEY (group_id) REFERENCES `groups`(group_id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
