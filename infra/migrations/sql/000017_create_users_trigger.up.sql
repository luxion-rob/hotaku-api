-- DELIMITER not required because migrations run with multiStatements=true
CREATE TRIGGER `trg_users_bu_set_deleted_flag`
BEFORE UPDATE ON `users`
FOR EACH ROW
BEGIN
    SET NEW.deleted_flag = (NEW.deleted_at IS NOT NULL);
END;
