-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inserted_feed_follows.*,
        users.name AS user_name,
        feeds.name AS feed_name
FROM inserted_feed_follows
INNER JOIN users ON users.id = inserted_feed_follows.user_id
INNER JOIN feeds ON feeds.id = inserted_feed_follows.feed_id;
        
-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user_name 
from feed_follows
INNER JOIN users ON users.id = user_id
INNER JOIN feeds ON feeds.id = feed_id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollowByUserAndURL :exec
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1 AND feed_id = $2;