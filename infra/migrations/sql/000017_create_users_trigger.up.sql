-- No need for DELIMITER when running via golang-migrate
CREATE TRIGGER `trg_users_bi_set_deleted_flag`
BEFORE INSERT ON `users`
FOR EACH ROW
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        SET NEW.deleted_flag = TRUE;
    ELSE
        SET NEW.deleted_flag = FALSE;
    END IF;
END;

CREATE TRIGGER `trg_users_bu_set_deleted_flag`
BEFORE UPDATE ON `users`
FOR EACH ROW
BEGIN
    IF NEW.deleted_at IS NOT NULL THEN
        SET NEW.deleted_flag = TRUE;
    ELSE
        SET NEW.deleted_flag = FALSE;
    END IF;
END;
