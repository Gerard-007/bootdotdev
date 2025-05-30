// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: posts.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, title, description, url, published, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, title, description, url, published, feed_id, created_at, updated_at
`

type CreatePostParams struct {
	ID          uuid.UUID
	Title       string
	Description string
	Url         string
	Published   time.Time
	FeedID      uuid.NullUUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Url,
		arg.Published,
		arg.FeedID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Url,
		&i.Published,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :one
DELETE FROM posts WHERE id = $1
RETURNING id, title, description, url, published, feed_id, created_at, updated_at
`

func (q *Queries) DeletePost(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, deletePost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Url,
		&i.Published,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPost = `-- name: GetPost :one
SELECT id, title, description, url, published, feed_id, created_at, updated_at FROM posts WHERE id = $1
`

func (q *Queries) GetPost(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Url,
		&i.Published,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPostsForUser = `-- name: GetPostsForUser :many
SELECT posts.id, posts.title, posts.description, posts.url, posts.published, posts.feed_id, posts.created_at, posts.updated_at
FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published DESC
LIMIT $2
`

type GetPostsForUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

func (q *Queries) GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsForUser, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Url,
			&i.Published,
			&i.FeedID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET title = $2, description = $3, url = $4, published = $5, feed_id = $6, updated_at = NOW()
WHERE id = $1
RETURNING id, title, description, url, published, feed_id, created_at, updated_at
`

type UpdatePostParams struct {
	ID          uuid.UUID
	Title       string
	Description string
	Url         string
	Published   time.Time
	FeedID      uuid.NullUUID
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Url,
		arg.Published,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Url,
		&i.Published,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
