CREATE TABLE users_notifications (
    user_id VARCHAR(36) NOT NULL UNIQUE,
    notification_id VARCHAR(36) NOT NULL,
    sent_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    seen_at TIMESTAMP,
    PRIMARY KEY (user_id, notification_id),
    INDEX idx_users_notifications_user_id (user_id),
    INDEX idx_users_notifications_notification_id (notification_id),
    CONSTRAINT fk_users_notifications_users FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_users_notifications_notifications FOREIGN KEY (notification_id) REFERENCES notifications(notification_id) ON DELETE CASCADE,
    INDEX idx_users_notifications_unseen (user_id, seen_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;