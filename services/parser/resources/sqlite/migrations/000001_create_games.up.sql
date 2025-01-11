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
)