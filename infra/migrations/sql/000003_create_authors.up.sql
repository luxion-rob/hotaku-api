CREATE TABLE `authors` (
    author_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36),
    author_name VARCHAR(255),
    author_bio TEXT,
    primary KEY (author_id)
);
