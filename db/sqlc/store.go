package db

import (
	"context"
	"database/sql"
	"fmt"
)

//Store provides all function to execute db queries and transaction
type Store struct{
	*Queries
	db *sql.DB
}

//New store create a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:db,
		Queries: New(db),
	}
}

//execTx executes a function within database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx,nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err %v ", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

//TransferTxParams contain the input parameters of the tranfer transection
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

//TransferTxResult is a transfer result of a transaction
type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

//TransferTx perform a money transfer from one account to other
//It creates a tranfer record, add account entries, and update account's balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		return nil
	})

	return result, err
}    

