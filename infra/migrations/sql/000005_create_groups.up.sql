CREATE TABLE `groups` (
    group_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    group_name VARCHAR(255),
    primary KEY (group_id)
);