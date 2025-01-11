-- name: findGame :one
SELECT sqlc.embed(g), sqlc.embed(gi)  FROM games as g
    LEFT JOIN game_info as gi
    ON g.id = gi.game_id
    WHERE g.id = ? LIMIT 1;

-- name: createGame :exec
INSERT INTO games(id, name) VALUES (?, ?);

-- name: createGameInfo :exec
INSERT INTO game_info(game_id, image_url, initial_price, final_price, discount_percent) VALUES (?,?,?,?,?);

-- name: findUserGames :many
SELECT
    g.id AS game_id,
    g.name AS game_name,
    gi.image_url,
    gi.initial_price,
    gi.final_price,
    gi.discount_percent
FROM users_games AS ug
JOIN games AS g ON ug.game_id = g.id
LEFT JOIN game_info AS gi ON g.id = gi.game_id
WHERE ug.user_id = ?;