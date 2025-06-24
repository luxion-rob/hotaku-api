-- Sample data cho Hotaku API database (bảng đã đổi sang snake_case)

INSERT INTO roles (role_id, role_name)
VALUES
    (1, 'Admin'),
    (2, 'User');

INSERT INTO users (
    id,
    created_at,
    updated_at,
    deleted_at,
    email,
    password,
    name
)
VALUES
    (
        1,
        '2025-06-20 08:00:00',
        '2025-06-21 09:15:00',
        NULL,
        'alice@example.com',
        'hashed_pw_alice',
        'Alice Nguyen'
    ),
    (
        2,
        '2025-06-21 10:30:00',
        '2025-06-22 11:45:00',
        NULL,
        'bob@example.com',
        'hashed_pw_bob',
        'Bob Tran'
    ),
    (
        3,
        '2025-06-22 12:00:00',
        '2025-06-23 13:00:00',
        '2025-06-23 13:30:00',
        'charlie@example.com',
        'hashed_pw_charlie',
        'Charlie Le'
    );

INSERT INTO authors (author_id, external_id, author_name, author_bio)
VALUES
    (
        1,
        '33333333-3333-3333-3333-333333333333',
        'Kaori Yuki',
        'Author of fantasy manga series.'
    ),
    (
        2,
        '44444444-4444-4444-4444-444444444444',
        'Eiichiro Oda',
        'Creator of One Piece.'
    );

INSERT INTO categories (category_id, external_id, category_name)
VALUES
    (
        1,
        '55555555-5555-5555-5555-555555555555',
        'Fantasy'
    ),
    (
        2,
        '66666666-6666-6666-6666-666666666666',
        'Adventure'
    );

INSERT INTO `groups` (group_id, external_id, group_name)
VALUES
    (
        1,
        '77777777-7777-7777-7777-777777777777',
        'ScanLords'
    ),
    (
        2,
        '88888888-8888-8888-8888-888888888888',
        'MangaWorld'
    );

INSERT INTO mangas (
    manga_id,
    external_id,
    title,
    created_at,
    status,
    description
)
VALUES
    (
        1,
        '99999999-9999-9999-9999-999999999999',
        'Fairy Dream',
        '2025-01-01 00:00:00',
        'ongoing',
        'A magical adventure.'
    ),
    (
        2,
        'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
        'Ocean Pirates',
        '2025-02-15 00:00:00',
        'completed',
        'Pirate crew adventures.'
    );

INSERT INTO manga_authors (manga_id, author_id)
VALUES
    (1, 1),
    (1, 2),
    (2, 2);

INSERT INTO manga_categories (manga_id, category_id)
VALUES
    (1, 1),
    (2, 2);

INSERT INTO manga_chapters (
    chapter_id,
    external_id,
    manga_id,
    chapter_number,
    title,
    created_at
)
VALUES
    (
        1,
        'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
        1,
        1,
        'Chapter One',
        '2025-03-01 10:00:00'
    ),
    (
        2,
        'cccccccc-cccc-cccc-cccc-cccccccccccc',
        1,
        2,
        'Chapter Two',
        '2025-04-01 10:00:00'
    ),
    (
        3,
        'dddddddd-dddd-dddd-dddd-dddddddddddd',
        2,
        1,
        'First Voyage',
        '2025-03-10 11:00:00'
    );

INSERT INTO chapter_pages (
    page_id,
    external_id,
    chapter_id,
    image_url,
    page_number
)
VALUES
    (
        1,
        'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee',
        1,
        'https://example.com/fairy1.jpg',
        1
    ),
    (
        2,
        'ffffffff-ffff-ffff-ffff-ffffffffffff',
        1,
        'https://example.com/fairy2.jpg',
        2
    ),
    (
        3,
        '11112222-3333-4444-5555-666677778888',
        2,
        'https://example.com/fairy3.jpg',
        1
    ),
    (
        4,
        '99990000-aaaa-bbbb-cccc-ddddeeeeffff',
        3,
        'https://example.com/pirate1.jpg',
        1
    );

INSERT INTO manga_groups (manga_id, group_id)
VALUES
    (1, 1),
    (2, 2);

INSERT INTO notifications (notification_id, external_id, message, created_at)
VALUES
    (
        1,
        'abababab-abab-abab-abab-abababababab',
        'New chapter 2 of Fairy Dream is out!',
        '2025-05-01 12:00:00'
    ),
    (
        2,
        'cdcdcdcd-cdcd-cdcd-cdcd-cdcdcdcdcdcd',
        'Ocean Pirates completed!',
        '2025-05-15 15:30:00'
    );

INSERT INTO user_favorite_mangas (
    favorite_id,
    external_id,
    user_id,
    manga_id,
    added_at
)
VALUES
    (
        1,
        'edededed-eded-eded-eded-edededededed',
        1,
        1,
        '2025-05-02 09:00:00'
    ),
    (
        2,
        'fdfdfdfd-fdfd-fdfd-fdfd-fdfdfdfdfdfd',
        2,
        2,
        '2025-05-16 10:00:00'
    );

INSERT INTO user_manga_historys (
    history_id,
    external_id,
    user_id,
    manga_id,
    read_chapter_ids,
    read_at
)
VALUES
    (
        1,
        '12121212-3434-5656-7878-909090909090',
        1,
        1,
        '1,2',
        '2025-05-05 14:00:00'
    ),
    (
        2,
        'abab1212-cdcd-efef-3434-565656565656',
        2,
        2,
        '3',
        '2025-05-20 16:00:00'
    );

INSERT INTO user_notifications (user_id, notification_id)
VALUES
    (1, 1),
    (2, 2);
