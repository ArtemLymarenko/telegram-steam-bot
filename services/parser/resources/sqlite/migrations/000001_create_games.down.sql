-- Drop triggers first
DROP TRIGGER IF EXISTS games_ai;
DROP TRIGGER IF EXISTS games_ad;
DROP TRIGGER IF EXISTS games_au;

-- Drop indexes
DROP INDEX IF EXISTS idx_games_name;

-- Drop FTS virtual table
DROP TABLE IF EXISTS games_fts;

-- Drop tables in correct order (respect foreign keys)
DROP TABLE IF EXISTS game_info;
DROP TABLE IF EXISTS games;