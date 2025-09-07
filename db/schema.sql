CREATE TABLE pets
(
    id               TEXT PRIMARY KEY,
    name             TEXT    NOT NULL,
    breed            TEXT    NOT NULL,
    gender           TEXT    NOT NULL,
    price            NUMERIC NOT NULL,
    image_url        TEXT,
    likes            INTEGER NOT NULL,
    shop_name        TEXT    NOT NULL,
    shop_location    TEXT   NOT NULL,
    birth_date       DATE,
    reference_number TEXT   NOT NULL,
    tags             TEXT[]  NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 新規に作成する reservations テーブル
CREATE TABLE IF NOT EXISTS reservations
(
    -- 予約ごとに一意のIDを持たせる (UUID, SERIALなど)
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    -- ユーザを識別するID（外部の認証IDや社内システムIDなど任意）
    user_id        TEXT    NOT NULL,
    -- ユーザの氏名
    user_name      TEXT    NOT NULL,
    -- ユーザのメールアドレス
    email          TEXT    NOT NULL,
    -- 見学予定日時
    reservation_date_time TIMESTAMP NOT NULL,
    -- 予約ステータス pending, confirmed, cancelled
    status TEXT NOT NULL,
    -- 予約対象のペットID
    -- petsテーブルのidを参照 (FK)
    pet_id         TEXT    NOT NULL REFERENCES pets (id) ON DELETE CASCADE,
    -- 予約レコードが作られた日時
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- 予約レコードが更新された日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- お気に入りを管理するテーブル
CREATE TABLE IF NOT EXISTS favorites
(
    -- 予約ごとに一意のIDを持たせる (UUID, SERIALなど)
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    -- ユーザを識別するID（外部の認証IDや社内システムIDなど任意）
    user_id        TEXT    NOT NULL,
    -- 予約対象のペットID
    -- petsテーブルのidを参照 (FK)
    pet_id         TEXT    NOT NULL REFERENCES pets (id) ON DELETE CASCADE
);

-- 通知を管理するテーブル
CREATE TABLE IF NOT EXISTS notifications
(
    -- 通知ごとに一意のIDを持たせる
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    -- ユーザを識別するID
    user_id        TEXT    NOT NULL,
    -- 通知のタイトル
    title          TEXT    NOT NULL,
    -- 通知のメッセージ内容
    message        TEXT    NOT NULL,
    -- 既読状態
    is_read        BOOLEAN NOT NULL DEFAULT FALSE,
    -- 通知の種類
    type           TEXT    NOT NULL,
    -- 通知レコードが作られた日時
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- 通知レコードが更新された日時
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

---

-- Indexes for reservations table
CREATE INDEX idx_reservations_pet_id ON reservations(pet_id);
CREATE INDEX idx_reservations_user_id ON reservations(user_id);
CREATE INDEX idx_reservations_status ON reservations(status);
-- Indexes for favorites table
CREATE INDEX idx_favorites_user_id ON favorites(user_id);
-- Composite unique index to ensure one favorite per user-pet combination
CREATE UNIQUE INDEX idx_favorites_user_pet ON favorites(user_id, pet_id);
-- Indexes for notifications table
CREATE INDEX idx_notifications_user_id ON notifications(user_id);

---

INSERT INTO pets (
    id,
    name,
    breed,
    gender,
    price,
    image_url,
    likes,
    shop_name,
    shop_location,
    birth_date,
    reference_number,
    tags
) VALUES
-- 1
('1',
 'cute cat',
 'brown cat',
 'Male',
 360000,
 'https://images.unsplash.com/photo-1583083527882-4bee9aba2eea?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=777&q=80',
 10,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000001',
 '{"cute","famous","cool","new"}'
),

-- 2
('2',
 'サイベリアン',
 'Siberian cat',
 'Female',
 400000,
 'https://images.unsplash.com/photo-1586289883499-f11d28aaf52f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxleHBsb3JlLWZlZWR8OHx8fGVufDB8fHx8&auto=format&fit=crop&w=500&q=60',
 3,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000002',
 '{"cute","famous","cool","on-sale"}'
),

-- 3
('3',
 'Red cat',
 'red cat',
 'Male',
 240000,
 'https://images.unsplash.com/photo-1606491048802-8342506d6471?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxleHBsb3JlLWZlZWR8MTF8fHxlbnwwfHx8fA%3D%3D&auto=format&fit=crop&w=500&q=60',
 7,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000003',
 '{"cute","famous","cool"}'
),

-- 4
('4',
 'cute kitten',
 'white cat',
 'Female',
 550000,
 'https://images.unsplash.com/photo-1605450648855-63f9161b7ef7?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8ODZ8fGtpdHRlbnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60',
 12,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000004',
 '{"cute","famous","cool"}'
),

-- 5
('5',
 'Matcha',
 'Minuet',
 'Male',
 400000,
 'https://pbs.twimg.com/media/FaxOK5HUIAANRYo?format=jpg&name=4096x4096',
 5,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000005',
 '{"くりくりの目","きれいな毛並み","おてんば"}'
),

-- 6 (id: uma-chan)
('77',
 'uma-chan',
 'kage',
 'Female',
 49800000,
 'https://images.unsplash.com/photo-1557413606-2a63a06a1f1d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=500&q=60',
 15,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000006',
 '{"cute","famous","cool"}'
),

-- 7 (id: arai-san)
('100',
 'arai-san',
 'white cat',
 'Male',
 50000,
 'https://images.unsplash.com/photo-1601247387326-f8bcb5a234d4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=500&q=60',
 9,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000007',
 '{"cute","famous","cool"}'
),

-- 8 (id: 6)
('6',
 'cute kitten',
 'white cat',
 'Female',
 550000,
 'https://images.unsplash.com/photo-1597626133663-53df9633b799?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxleHBsb3JlLWZlZWR8MTV8fHxlbnwwfHx8fA%3D%3D&auto=format&fit=crop&w=500&q=60',
 2,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000008',
 '{"cute","famous","cool"}'
),

-- 9 (id: 7)
('7',
 'cute kitten',
 'white cat',
 'Male',
 550000,
 'https://images.unsplash.com/photo-1621238281284-d186cb6813fb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTh8fGtpdHRlbnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60',
 11,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000009',
 '{"cute","famous","cool"}'
),

-- 10 (id: 8)
('8',
 'cute kitten',
 'white cat',
 'Female',
 550000,
 'https://images.unsplash.com/photo-1557166984-b00337652c94?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTl8fGtpdHRlbnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60',
 4,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000010',
 '{"cute","famous","cool"}'
),

-- 11 (id: 9)
('9',
 'cute kitten',
 'white cat',
 'Male',
 550000,
 'https://images.unsplash.com/flagged/photo-1557427161-4701a0fa2f42?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mjh8fGtpdHRlbnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60',
 8,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000011',
 '{"cute","famous","cool"}'
),

-- 12 (id: 10)
('10',
 'cute kitten',
 'white cat',
 'Female',
 550000,
 'https://images.unsplash.com/photo-1582797493098-23d8d0cc6769?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8NjZ8fGtpdHRlbnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60',
 6,
 'uma-arai shop 2nd',
 'Kanagawa',
 '2023-10-14',
 '0000012',
 '{"cute","famous","cool"}'
);

-- Notification test data
INSERT INTO notifications (
    user_id,
    title,
    message,
    is_read,
    type
) VALUES
('user1', '新しい子が増えました', '可愛い「あらいさん」が新しく増えました。ぜひチェックしてください！', false, 'new_pet'),
('user1', 'お気に入りペットの価格変更', 'お気に入り登録しているペットの価格が変更されました。', false, 'price_change'),
('user2', '予約確認', 'お目当ての子：「うまちゃん」。', true, 'reservation'),
('user1', 'キャンペーン情報', '今週末限定のキャンペーンが開始されました！', false, 'campaign');
