-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetNotFollowedFeeds :many
SELECT feeds.*
FROM feeds
LEFT JOIN feed_follows ON
	feeds.id = feed_follows.feed_id
	AND feed_follows.user_id = $1
WHERE feed_follows.feed_id IS NULL;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds 
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET 
    last_fetched_at = timezone('utc', now()),
    updated_at = timezone('utc', now())
WHERE id = $1
RETURNING *;