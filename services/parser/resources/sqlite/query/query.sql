-- name: FindGame :one
SELECT sqlc.embed(g), sqlc.embed(gi)  FROM games as g
    LEFT JOIN game_info as gi
    ON g.id = gi.game_id
    WHERE g.id = ? LIMIT 1;

-- name: CreateGame :exec
INSERT INTO games(id, name) VALUES (?, ?);

-- name: CreateGameInfo :exec
INSERT INTO game_info(game_id, image_url, initial_price, final_price, discount_percent) VALUES (?,?,?,?,?);