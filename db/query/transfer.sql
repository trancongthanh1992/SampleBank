-- name: GetTransfer :one
SELECT * FROM transfers 
WHERE id = $1;

-- name: ListTransfer :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $1
WHERE id = $2
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;