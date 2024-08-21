-- name: GetPostsForUser :many
SELECT * FROM posts AS p
JOIN feed_follows AS f 
ON p.feed_id = f.feed_id
WHERE f.user_id = $1
ORDER BY p.published_at DESC
LIMIT $2;