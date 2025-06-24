CREATE TABLE `manga_groups` (
        manga_id INT NOT NULL,
        group_id INT NOT NULL,
        primary KEY (manga_id, group_id),
        foreign KEY (manga_id) REFERENCES `mangas` (manga_id) ON DELETE CASCADE,
        foreign KEY (group_id) REFERENCES `groups` (group_id) ON DELETE CASCADE
    );