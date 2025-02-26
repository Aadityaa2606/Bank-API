-- name: CreateSession :one
INSERT INTO sessions (
    id, username, refresh_token, is_revoked, expires_at
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;

-- name: RevokeSession :exec
UPDATE sessions
SET is_revoked = true
WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1;