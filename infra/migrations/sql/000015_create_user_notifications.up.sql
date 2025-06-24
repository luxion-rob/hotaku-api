CREATE TABLE `user_notifications` (
        user_id BIGINT UNSIGNED NOT NULL,
        notification_id INT NOT NULL,
        primary KEY (user_id, notification_id),
        foreign KEY (user_id) REFERENCES `users` (id) ON DELETE CASCADE,
        foreign KEY (notification_id) REFERENCES `notifications` (notification_id) ON DELETE CASCADE
    );