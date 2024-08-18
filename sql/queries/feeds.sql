-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1,$2,$3,$4,$5,$6) 
RETURNING *;

-- name: IsFeedExists :one
SELECT EXISTS(SELECT 1 FROM feeds WHERE id = $1);

-- name: GetAllFeeds :many
SELECT * FROM feeds;