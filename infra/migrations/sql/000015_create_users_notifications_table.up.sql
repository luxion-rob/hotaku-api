CREATE TABLE `users_notifications` (
    user_id CHAR(36) NOT NULL,
    notification_id CHAR(36) NOT NULL,
    sent_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    seen_at TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (user_id, notification_id),
    CONSTRAINT fk_users_notifications_users
        FOREIGN KEY (user_id) REFERENCES users(user_id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_users_notifications_notifications
        FOREIGN KEY (notification_id) REFERENCES notifications(notification_id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    INDEX idx_users_notifications_user_id_seen_at (user_id, seen_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;