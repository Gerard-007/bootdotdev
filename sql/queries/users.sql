-- name: CreateUser :one
INSERT INTO users (id, email, password, username, created_at, updated_at, api_key) 
VALUES ($1, $2, $3, $4, $5, $6, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;


-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;