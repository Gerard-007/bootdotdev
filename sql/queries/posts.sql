-- name: CreatePost :one
INSERT INTO posts (id, title, description, url, published, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.*
FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published DESC
LIMIT $2;

-- name: GetPost :one
SELECT * FROM posts WHERE id = $1;

-- name: UpdatePost :one
UPDATE posts
SET title = $2, description = $3, url = $4, published = $5, feed_id = $6, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePost :one
DELETE FROM posts WHERE id = $1
RETURNING *;