CREATE TABLE `mangas` (
        manga_id INT NOT NULL AUTO_INCREMENT,
        external_id CHAR(36),
        title VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        status VARCHAR(100),
        description TEXT,
        primary KEY (manga_id)
    );