CREATE TABLE IF NOT EXISTS users_games (
    game_id          INTEGER NOT NULL,
    user_id 		 INTEGER NOT NULL,

    CONSTRAINT games_game_id_fk
        FOREIGN KEY (game_id)
            REFERENCES games (id)
            ON DELETE CASCADE
            ON UPDATE CASCADE,
   UNIQUE (game_id, user_id)
);