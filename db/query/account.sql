CREATE TABLE accounts (
  id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  owner varchar(255) NOT NULL,
  balance bigint NOT NULL,
  currency varchar(3) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- name: CreateAccounts :execresult
INSERT INTO accounts (
  owner,balance,currency
) VALUES (
  ?, ?, ?
);

-- name: UpdateAccount :execresult
UPDATE accounts 
SET owner = ?,balance = ?,currency = ?
WHERE id = ?;

-- name: GetLastAccount :one
SELECT id,owner,balance,currency,created_at FROM accounts
ORDER BY id DESC 
LIMIT 1;

-- name: GetAccountByID :one
SELECT id,owner,balance,currency,created_at FROM accounts
WHERE id = ?;


-- name: GetAccountByIDForUpdate :one
SELECT id,owner,balance,currency,created_at FROM accounts
WHERE id = ?
FOR UPDATE;

-- name: ListAccounts :many
SELECT id,owner,balance,currency,created_at FROM accounts
ORDER BY id
LIMIT ? 
OFFSET ?;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = ?;
