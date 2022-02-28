CREATE TABLE transfers (
  id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  from_account_id bigint NOT NULL,
  to_account_id bigint NOT NULL,
  amount bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- name: CreateTransfer :execresult
INSERT INTO transfers (
  from_account_id,to_account_id,amount
) VALUES (
  ?, ?, ?
);

-- name: GetLastTransfer :one
SELECT id,from_account_id,to_account_id,amount,created_at FROM transfers
ORDER BY id DESC 
LIMIT 1;

-- name: GetTransferByID :one
SELECT id,from_account_id,to_account_id,amount,created_at FROM transfers
WHERE id = ?;

-- name: ListTransfers :many
SELECT id,from_account_id,to_account_id,amount,created_at FROM transfers
ORDER BY id;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = ?;