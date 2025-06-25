CREATE TABLE notifications (
    notification_id INT NOT NULL AUTO_INCREMENT,
    external_id CHAR(36) NOT NULL UNIQUE,
    message VARCHAR(500) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (notification_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;