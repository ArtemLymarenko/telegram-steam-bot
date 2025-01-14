-- name: findGame :one
SELECT sqlc.embed(g), sqlc.embed(gi)  FROM games as g
    LEFT JOIN game_info as gi
    ON g.id = gi.game_id
    WHERE g.id = ? LIMIT 1;

-- name: createGame :exec
INSERT INTO games(id, name) VALUES (?, ?);

-- name: createGameInfo :exec
INSERT INTO game_info(game_id, url, image_url, initial_price, final_price, discount_percent) VALUES (?, ?, ?, ?, ?, ?);

-- name: findUserGames :many
SELECT sqlc.embed(g), sqlc.embed(gi)
FROM users_games AS ug
JOIN games AS g ON ug.game_id = g.id
LEFT JOIN game_info AS gi ON g.id = gi.game_id
WHERE ug.user_id = ?;

-- name: addUserGame :exec
INSERT INTO users_games(user_id, game_id) VALUES (?, ?);

-- name: deleteGameById :exec
DELETE FROM games WHERE id = ?;

-- name: searchGamesByName :many
SELECT sqlc.embed(g), sqlc.embed(gi)
FROM games_fts AS fts
JOIN games AS g ON fts.rowid = g.id
LEFT JOIN game_info AS gi ON g.id = gi.game_id
WHERE fts.name  MATCH ?;

-- name: deleteUserGame :one
DELETE FROM users_games WHERE user_id = ? AND game_id = ? RETURNING game_id;