CREATE TABLE IF NOT EXISTS users_games (
    game_id          INTEGER NOT NULL,
    user_id 		 INTEGER NOT NULL,

    FOREIGN KEY (game_id) REFERENCES games(id)
       ON DELETE CASCADE
       ON UPDATE CASCADE
);