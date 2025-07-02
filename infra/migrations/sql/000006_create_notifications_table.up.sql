CREATE TABLE notifications (
    notification_id CHAR(36) NOT NULL,
    external_id CHAR(36) NOT NULL,
    message VARCHAR(500) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (notification_id),
    INDEX idx_notifications_created_at (created_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;