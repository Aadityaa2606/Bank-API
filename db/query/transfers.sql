-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id, to_account_id, amount
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetTransferByID :one
SELECT * FROM transfers
WHERE id = $1;

-- name: GetTransferByFromAccountID :many
SELECT * FROM transfers
WHERE from_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: GetTransferByToAccountID :many
SELECT * FROM transfers
WHERE to_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $1
WHERE id = $2
RETURNING *;

-- name: DeleteTransfer :one
DELETE FROM transfers
WHERE id = $1
RETURNING *;