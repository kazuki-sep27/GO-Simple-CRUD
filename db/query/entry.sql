CREATE TABLE entries (
  id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  account_id bigint NOT NULL,
  amount bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- name: CreateEntry :execresult
INSERT INTO entries (
  account_id,amount
) VALUES (
  ?, ?
);

-- name: GetLastEntry :one
SELECT id,account_id,amount,created_at FROM entries
ORDER BY id DESC 
LIMIT 1;

-- name: GetEntryByID :one
SELECT id,account_id,amount,created_at FROM entries
WHERE id = ?;

-- name: ListEntries :many
SELECT id,account_id,amount,created_at FROM entries
ORDER BY created_at DESC;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = ?;