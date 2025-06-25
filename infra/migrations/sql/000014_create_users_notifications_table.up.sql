CREATE TABLE users_notifications (
    user_id BIGINT NOT NULL,
    notification_id INT NOT NULL,
    sent_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    seen_at TIMESTAMP,
    PRIMARY KEY (user_id, notification_id),
    INDEX idx_users_notifications_user_id (user_id),
    INDEX idx_users_notifications_notification_id (notification_id),
    CONSTRAINT fk_users_notifications_users FOREIGN KEY (user_id) REFERENCES users(user_id),
    CONSTRAINT fk_users_notifications_notifications FOREIGN KEY (notification_id) REFERENCES notifications(notification_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;