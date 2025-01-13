-- Base tables for storing games and their information
CREATE TABLE IF NOT EXISTS games (
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS game_info (
    game_id          INTEGER PRIMARY KEY,
    url              TEXT,
    image_url        TEXT,
    initial_price    REAL,
    final_price      REAL,
    discount_percent REAL,

    CONSTRAINT game_info_game_id_fk
        FOREIGN KEY (game_id) REFERENCES games (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- FTS table for fast text search
CREATE VIRTUAL TABLE IF NOT EXISTS games_fts USING fts5(
    name,
    content='games',
    content_rowid='id'
);

-- Triggers to keep FTS index in sync with games table
CREATE TRIGGER IF NOT EXISTS games_ai AFTER INSERT ON games BEGIN
    INSERT INTO games_fts(rowid, name) VALUES (new.id, new.name);
END;

CREATE TRIGGER IF NOT EXISTS games_ad AFTER DELETE ON games BEGIN
    DELETE FROM games_fts WHERE rowid = old.id;
END;

CREATE TRIGGER IF NOT EXISTS games_au AFTER UPDATE ON games BEGIN
    REPLACE INTO games_fts(rowid, name) VALUES (new.id, new.name);
END;

-- Create index for better search performance
CREATE INDEX IF NOT EXISTS idx_games_name ON games(name);