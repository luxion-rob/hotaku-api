-- =============================================
-- Sample data for roles
-- =============================================
INSERT INTO `roles` (`role_id`, `role_name`) VALUES
  ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Admin'),
  ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'User'),
  ('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Guest');

-- =============================================
-- Sample data for users
-- =============================================
INSERT INTO `users` (`user_id`, `role_id`, `email`, `password`, `name`) VALUES
  ('00000001-0001-0001-0001-000000000001', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'admin@example.com', 'adminpass',    'Administrator'),
  ('00000002-0002-0002-0002-000000000002', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'john.doe@example.com', 'userpass', 'John Doe'),
  ('00000003-0003-0003-0003-000000000003', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 'jane.smith@example.com', 'guestpass','Jane Smith');

-- =============================================
-- Sample data for authors
-- =============================================
INSERT INTO `authors` (`author_id`, `external_id`, `author_name`, `author_bio`) VALUES
  ('11111111-1111-1111-1111-111111111111', '22222222-2222-2222-2222-222222222222', 'Eiichiro Oda',    'Tác giả của One Piece'),
  ('33333333-3333-3333-3333-333333333333', '44444444-4444-4444-4444-444444444444', 'Masashi Kishimoto','Tác giả của Naruto'),
  ('55555555-5555-5555-5555-555555555555', '66666666-6666-6666-6666-666666666666', 'Tite Kubo',        'Tác giả của Bleach');

-- =============================================
-- Sample data for groups
-- =============================================
INSERT INTO `groups` (`group_id`, `external_id`, `group_name`) VALUES
  ('77777777-7777-7777-7777-777777777777', '88888888-8888-8888-8888-888888888888', 'Funimation'),
  ('99999999-9999-9999-9999-999999999999', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaa0000', 'Crunchyroll'),
  ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbb0000', 'cccccccc-cccc-cccc-cccc-cccccccc0000', 'Viz Media');

-- =============================================
-- Sample data for categories
-- =============================================
INSERT INTO `categories` (`category_id`, `external_id`, `category_name`) VALUES
  ('dddddddd-dddd-dddd-dddd-dddddddddddd', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'Shonen'),
  ('ffffffff-ffff-ffff-ffff-ffffffffffff', '00000000-0000-0000-0000-000000000000', 'Seinen'),
  ('12121212-1212-1212-1212-121212121212', '34343434-3434-3434-3434-343434343434', 'Comedy');

-- =============================================
-- Sample data for notifications
-- =============================================
INSERT INTO `notifications` (`notification_id`, `external_id`, `message`) VALUES
  ('88888888-8888-8888-8888-888888888888', '99999999-9999-9999-9999-999999999999', 'Welcome to the platform!'),
  ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaa1111', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'New manga chapter available!');

-- =============================================
-- Sample data for manga_statuses
-- =============================================
INSERT INTO `manga_statuses` (`status_id`, `status_name`) VALUES
  (1, 'Ongoing'),
  (2, 'Completed'),
  (3, 'Hiatus');

-- =============================================
-- Sample data for mangas
-- =============================================
INSERT INTO `mangas` (`manga_id`, `external_id`, `status_id`, `title`, `description`) VALUES
  ('cccccccc-0001-0001-0001-000000000001', 'dddddddd-0001-0001-0001-000000000001', 1, 'One Piece', 'Cuộc phiêu lưu của băng Mũ Rơm'),
  ('cccccccc-0002-0002-0002-000000000002', 'dddddddd-0002-0002-0002-000000000002', 2, 'Naruto',    'Hành trình của Uzumaki Naruto'),
  ('cccccccc-0003-0003-0003-000000000003', 'dddddddd-0003-0003-0003-000000000003', 1, 'Bleach',    'Cuộc chiến với Hollow');

-- =============================================
-- Sample data for manga_chapters
-- =============================================
INSERT INTO `manga_chapters` (`chapter_id`, `external_id`, `manga_id`, `chapter_number`, `title`) VALUES
  ('eeeeeeee-1001-1001-1001-000000010001', 'ffffffff-1001-1001-1001-000000010001', 'cccccccc-0001-0001-0001-000000000001', 1, 'Romance Dawn'),
  ('eeeeeeee-1001-1001-1001-000000010002', 'ffffffff-1001-1001-1001-000000010002', 'cccccccc-0001-0001-0001-000000000001', 2, 'They Call Him Strawhat Luffy'),
  ('eeeeeeee-2002-2002-2002-000000020001', 'ffffffff-2002-2002-2002-000000020001', 'cccccccc-0002-0002-0002-000000000002', 1, 'Uzumaki Naruto'),
  ('eeeeeeee-3003-3003-3003-000000030001', 'ffffffff-3003-3003-3003-000000030001', 'cccccccc-0003-0003-0003-000000000003', 1, 'Death & Strawberry');

-- =============================================
-- Sample data for chapter_pages
-- =============================================
INSERT INTO `chapter_pages` (`page_id`, `external_id`, `chapter_id`, `page_number`, `image_url`) VALUES
  ('11111111-2001-2001-2001-000000100001', '22222222-2001-2001-2001-000000100001', 'eeeeeeee-1001-1001-1001-000000010001', 1, 'http://example.com/op_ch1_p1.jpg'),
  ('11111111-2001-2001-2001-000000100002', '22222222-2001-2001-2001-000000100002', 'eeeeeeee-1001-1001-1001-000000010001', 2, 'http://example.com/op_ch1_p2.jpg'),
  ('11111111-2002-2002-2002-000000200001', '22222222-2002-2002-2002-000000200001', 'eeeeeeee-2002-2002-2002-000000020001', 1, 'http://example.com/naruto_ch1_p1.jpg'),
  ('11111111-3003-3003-3003-000000300001', '22222222-3003-3003-3003-000000300001', 'eeeeeeee-3003-3003-3003-000000030001', 1, 'http://example.com/bleach_ch1_p1.jpg');

-- =============================================
-- Sample data for mangas_authors
-- =============================================
INSERT INTO `mangas_authors` (`manga_id`, `author_id`) VALUES
  ('cccccccc-0001-0001-0001-000000000001', '11111111-1111-1111-1111-111111111111'),
  ('cccccccc-0002-0002-0002-000000000002', '33333333-3333-3333-3333-333333333333'),
  ('cccccccc-0003-0003-0003-000000000003', '55555555-5555-5555-5555-555555555555');

-- =============================================
-- Sample data for mangas_categories
-- =============================================
INSERT INTO `mangas_categories` (`manga_id`, `category_id`) VALUES
  ('cccccccc-0001-0001-0001-000000000001', 'dddddddd-dddd-dddd-dddd-dddddddddddd'),
  ('cccccccc-0002-0002-0002-000000000002', 'dddddddd-dddd-dddd-dddd-dddddddddddd'),
  ('cccccccc-0003-0003-0003-000000000003', 'dddddddd-dddd-dddd-dddd-dddddddddddd');

-- =============================================
-- Sample data for mangas_groups
-- =============================================
INSERT INTO `mangas_groups` (`manga_id`, `group_id`) VALUES
  ('cccccccc-0001-0001-0001-000000000001', '77777777-7777-7777-7777-777777777777'),
  ('cccccccc-0002-0002-0002-000000000002', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbb0000'),
  ('cccccccc-0003-0003-0003-000000000003', '99999999-9999-9999-9999-999999999999');

-- =============================================
-- Sample data for user_favorite_mangas
-- =============================================
INSERT INTO `user_favorite_mangas` (`favorite_id`, `external_id`, `user_id`, `manga_id`) VALUES
  ('abcabcab-000f-000f-000f-00000000f001', 'defdefde-000f-000f-000f-00000000f001', '00000002-0002-0002-0002-000000000002', 'cccccccc-0001-0001-0001-000000000001'),
  ('abcabcab-000f-000f-000f-00000000f002', 'defdefde-000f-000f-000f-00000000f002', '00000002-0002-0002-0002-000000000002', 'cccccccc-0002-0002-0002-000000000002');

-- =============================================
-- Sample data for users_notifications
-- =============================================
INSERT INTO `users_notifications` (`user_id`, `notification_id`) VALUES
  ('00000002-0002-0002-0002-000000000002', '88888888-8888-8888-8888-888888888888'),
  ('00000003-0003-0003-0003-000000000003', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaa1111');

-- =============================================
-- Sample data for user_read_chapters
-- =============================================
INSERT INTO `user_read_chapters` (`user_id`, `chapter_id`) VALUES
  ('00000002-0002-0002-0002-000000000002', 'eeeeeeee-1001-1001-1001-000000010001'),
  ('00000003-0003-0003-0003-000000000003', 'eeeeeeee-3003-3003-3003-000000030001');

