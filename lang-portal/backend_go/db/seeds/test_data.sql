-- Test data for Words table
INSERT INTO words (japanese, romaji, english, parts) VALUES
('こんにちは', 'konnichiwa', 'hello', '{"type": "greeting", "formality": "neutral"}'),
('ありがとう', 'arigatou', 'thank you', '{"type": "expression", "formality": "neutral"}'),
('猫', 'neko', 'cat', '{"type": "noun", "category": "animals"}'),
('食べる', 'taberu', 'to eat', '{"type": "verb", "conjugation": "ichidan"}'),
('大きい', 'ookii', 'big', '{"type": "adjective", "category": "size"}');

-- Test data for Groups table
INSERT INTO groups (name, description) VALUES
('Basic Greetings', 'Essential Japanese greetings and expressions'),
('Common Animals', 'Basic animal vocabulary for beginners'),
('Daily Verbs', 'Common everyday action words');

-- Test data for word_groups table
INSERT INTO word_groups (word_id, group_id) VALUES
(1, 1),  -- Basic Greetings - hello
(2, 1),  -- Basic Greetings - thank you
(3, 2),  -- Common Animals - cat
(4, 3);  -- Daily Verbs - to eat

-- Test data for study_activities
INSERT INTO study_activities (group_id, activity_type, created_at) VALUES
(1, 'flashcard', CURRENT_TIMESTAMP),
(1, 'quiz', CURRENT_TIMESTAMP),
(2, 'flashcard', CURRENT_TIMESTAMP);

-- Test data for study_sessions
INSERT INTO study_sessions (start_time, end_time, score, status, study_activity_id) VALUES
(CURRENT_TIMESTAMP, datetime(CURRENT_TIMESTAMP, '+5 minutes'), 85.0, 'completed', 1),
(CURRENT_TIMESTAMP, NULL, NULL, 'active', 2),
(datetime(CURRENT_TIMESTAMP, '-1 day'), datetime(CURRENT_TIMESTAMP, '-1 day', '+4 minutes'), 95.0, 'completed', 3);

-- Test data for word_review_items
INSERT INTO word_review_items (session_id, word_id, is_correct, reviewed_at) VALUES
(1, 1, 1, CURRENT_TIMESTAMP),  -- Correct review of こんにちは
(1, 2, 1, CURRENT_TIMESTAMP),  -- Correct review of ありがとう
(2, 3, 0, CURRENT_TIMESTAMP),  -- Incorrect review of 猫
(3, 4, 1, datetime(CURRENT_TIMESTAMP, '-1 day')); -- Correct review of 食べる from yesterday 