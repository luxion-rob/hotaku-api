CREATE TABLE `manga_statuses` (
    status_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    status_name VARCHAR(50) NOT NULL UNIQUE,
    PRIMARY KEY (status_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;