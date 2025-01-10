-- name: FindGame :one
SELECT g.id, g.name FROM games as g WHERE g.id = ? LIMIT 1;

-- name: CreateGame :one
INSERT INTO games(id, name) VALUES (?, ?) RETURNING id, name;