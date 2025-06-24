CREATE TABLE `notifications` (
        notification_id INT NOT NULL AUTO_INCREMENT,
        external_id CHAR(36),
        message VARCHAR(500),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        primary KEY (notification_id)
    );