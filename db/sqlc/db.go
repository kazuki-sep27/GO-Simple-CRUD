// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createAccountsStmt, err = db.PrepareContext(ctx, createAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccounts: %w", err)
	}
	if q.createEntryStmt, err = db.PrepareContext(ctx, createEntry); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEntry: %w", err)
	}
	if q.createTransferStmt, err = db.PrepareContext(ctx, createTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTransfer: %w", err)
	}
	if q.deleteAccountStmt, err = db.PrepareContext(ctx, deleteAccount); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAccount: %w", err)
	}
	if q.deleteEntryStmt, err = db.PrepareContext(ctx, deleteEntry); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteEntry: %w", err)
	}
	if q.deleteTransferStmt, err = db.PrepareContext(ctx, deleteTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTransfer: %w", err)
	}
	if q.getAccountByIDStmt, err = db.PrepareContext(ctx, getAccountByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccountByID: %w", err)
	}
	if q.getAccountByIDForUpdateStmt, err = db.PrepareContext(ctx, getAccountByIDForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccountByIDForUpdate: %w", err)
	}
	if q.getEntryByIDStmt, err = db.PrepareContext(ctx, getEntryByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetEntryByID: %w", err)
	}
	if q.getLastAccountStmt, err = db.PrepareContext(ctx, getLastAccount); err != nil {
		return nil, fmt.Errorf("error preparing query GetLastAccount: %w", err)
	}
	if q.getLastEntryStmt, err = db.PrepareContext(ctx, getLastEntry); err != nil {
		return nil, fmt.Errorf("error preparing query GetLastEntry: %w", err)
	}
	if q.getLastTransferStmt, err = db.PrepareContext(ctx, getLastTransfer); err != nil {
		return nil, fmt.Errorf("error preparing query GetLastTransfer: %w", err)
	}
	if q.getTransferByIDStmt, err = db.PrepareContext(ctx, getTransferByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransferByID: %w", err)
	}
	if q.listAccountsStmt, err = db.PrepareContext(ctx, listAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query ListAccounts: %w", err)
	}
	if q.listEntriesStmt, err = db.PrepareContext(ctx, listEntries); err != nil {
		return nil, fmt.Errorf("error preparing query ListEntries: %w", err)
	}
	if q.listTransfersStmt, err = db.PrepareContext(ctx, listTransfers); err != nil {
		return nil, fmt.Errorf("error preparing query ListTransfers: %w", err)
	}
	if q.updateAccountStmt, err = db.PrepareContext(ctx, updateAccount); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAccount: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createAccountsStmt != nil {
		if cerr := q.createAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccountsStmt: %w", cerr)
		}
	}
	if q.createEntryStmt != nil {
		if cerr := q.createEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createEntryStmt: %w", cerr)
		}
	}
	if q.createTransferStmt != nil {
		if cerr := q.createTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTransferStmt: %w", cerr)
		}
	}
	if q.deleteAccountStmt != nil {
		if cerr := q.deleteAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAccountStmt: %w", cerr)
		}
	}
	if q.deleteEntryStmt != nil {
		if cerr := q.deleteEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteEntryStmt: %w", cerr)
		}
	}
	if q.deleteTransferStmt != nil {
		if cerr := q.deleteTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTransferStmt: %w", cerr)
		}
	}
	if q.getAccountByIDStmt != nil {
		if cerr := q.getAccountByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountByIDStmt: %w", cerr)
		}
	}
	if q.getAccountByIDForUpdateStmt != nil {
		if cerr := q.getAccountByIDForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccountByIDForUpdateStmt: %w", cerr)
		}
	}
	if q.getEntryByIDStmt != nil {
		if cerr := q.getEntryByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEntryByIDStmt: %w", cerr)
		}
	}
	if q.getLastAccountStmt != nil {
		if cerr := q.getLastAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLastAccountStmt: %w", cerr)
		}
	}
	if q.getLastEntryStmt != nil {
		if cerr := q.getLastEntryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLastEntryStmt: %w", cerr)
		}
	}
	if q.getLastTransferStmt != nil {
		if cerr := q.getLastTransferStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLastTransferStmt: %w", cerr)
		}
	}
	if q.getTransferByIDStmt != nil {
		if cerr := q.getTransferByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransferByIDStmt: %w", cerr)
		}
	}
	if q.listAccountsStmt != nil {
		if cerr := q.listAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAccountsStmt: %w", cerr)
		}
	}
	if q.listEntriesStmt != nil {
		if cerr := q.listEntriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listEntriesStmt: %w", cerr)
		}
	}
	if q.listTransfersStmt != nil {
		if cerr := q.listTransfersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTransfersStmt: %w", cerr)
		}
	}
	if q.updateAccountStmt != nil {
		if cerr := q.updateAccountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAccountStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                          DBTX
	tx                          *sql.Tx
	createAccountsStmt          *sql.Stmt
	createEntryStmt             *sql.Stmt
	createTransferStmt          *sql.Stmt
	deleteAccountStmt           *sql.Stmt
	deleteEntryStmt             *sql.Stmt
	deleteTransferStmt          *sql.Stmt
	getAccountByIDStmt          *sql.Stmt
	getAccountByIDForUpdateStmt *sql.Stmt
	getEntryByIDStmt            *sql.Stmt
	getLastAccountStmt          *sql.Stmt
	getLastEntryStmt            *sql.Stmt
	getLastTransferStmt         *sql.Stmt
	getTransferByIDStmt         *sql.Stmt
	listAccountsStmt            *sql.Stmt
	listEntriesStmt             *sql.Stmt
	listTransfersStmt           *sql.Stmt
	updateAccountStmt           *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                          tx,
		tx:                          tx,
		createAccountsStmt:          q.createAccountsStmt,
		createEntryStmt:             q.createEntryStmt,
		createTransferStmt:          q.createTransferStmt,
		deleteAccountStmt:           q.deleteAccountStmt,
		deleteEntryStmt:             q.deleteEntryStmt,
		deleteTransferStmt:          q.deleteTransferStmt,
		getAccountByIDStmt:          q.getAccountByIDStmt,
		getAccountByIDForUpdateStmt: q.getAccountByIDForUpdateStmt,
		getEntryByIDStmt:            q.getEntryByIDStmt,
		getLastAccountStmt:          q.getLastAccountStmt,
		getLastEntryStmt:            q.getLastEntryStmt,
		getLastTransferStmt:         q.getLastTransferStmt,
		getTransferByIDStmt:         q.getTransferByIDStmt,
		listAccountsStmt:            q.listAccountsStmt,
		listEntriesStmt:             q.listEntriesStmt,
		listTransfersStmt:           q.listTransfersStmt,
		updateAccountStmt:           q.updateAccountStmt,
	}
}
