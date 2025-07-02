CREATE TABLE authors (
    author_id CHAR(36) NOT NULL,
    external_id CHAR(36) NOT NULL,
    author_name VARCHAR(50) NOT NULL,
    author_bio TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (author_id),
    UNIQUE KEY uq_authors_author_name (author_name)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;