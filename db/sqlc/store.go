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

//TransferResultTx is the result of the transfer transaction
type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

var txKey = struct{}{}

//TransferTx perform a money transfer from one account to other
//It creates a tranfer record, add account entries, and update account's balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		//txName := ctx.Value(txKey)

		_ ,err = q.query(ctx,nil,"SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED")
		if err != nil {
			return err
		}

		//fmt.Println(txName, "Create Transfer")

		_ ,err = q.CreateTransfer(ctx , CreateTransferParams {
			FromAccountID : arg.FromAccountID,
			ToAccountID : arg.ToAccountID,
			Amount : arg.Amount,
		})
		
		if err != nil {
			return err
		}

		result.Transfer,_ = q.GetLastTransfer(ctx)

		//fmt.Println(txName, "Create Entry 1")
		_, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry,_ = q.GetLastEntry(ctx)

		//fmt.Println(txName, "Create Entry 2")
		_, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry,_ = q.GetLastEntry(ctx)

		//get account -> update account's balance
		//update from account
		//fmt.Println(txName, "Get Account 1 before update")
		account1, err := q.GetAccountByIDForUpdate(ctx,arg.FromAccountID)
		if err != nil {
			return err
		}

		//fmt.Println(txName, "Update Account 1")
		_, err = q.UpdateAccount(ctx,UpdateAccountParams{
			ID: account1.ID,
			Balance: account1.Balance - arg.Amount,
		})
		if err != nil {
			return err
		}

		//fmt.Println(txName, "Get Account 1 after update")
		result.FromAccount, err = q.GetAccountByIDForUpdate(ctx,account1.ID)
		if err != nil {
			return err
		} 

		//update to account
		//fmt.Println(txName, "Get Account 2 before update")
		account2, err := q.GetAccountByIDForUpdate(ctx,arg.ToAccountID)
		if err != nil {
			return err
		}

		//fmt.Println(txName, "Update Account 2")
		_, err = q.UpdateAccount(ctx,UpdateAccountParams{
			ID: account2.ID,
			Balance: account2.Balance + arg.Amount,
		})
		if err != nil {
			return err
		}

		//fmt.Println(txName, "Get Account 2 after update")
		result.ToAccount, err = q.GetAccountByIDForUpdate(ctx,account2.ID)
		if err != nil {
			return err
		} 

		return nil
	})

	return result, err
}    

