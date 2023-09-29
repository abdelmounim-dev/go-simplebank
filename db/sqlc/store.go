package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions.
type Store struct {
  *Queries
  db *sql.DB
}

// NewStore creates a new Store.
func NewStore(db *sql.DB) *Store {
  return &Store{
    db:      db,
    Queries: New(db),
  }
}

// execTx executes a function within a database transaction.
// The provided function receives a transaction NewStore.
// If the function returns an error, the transaction is rolled back.
// If the function returns nil, the transaction is committed.
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
  tx, err := store.db.BeginTx(ctx, nil)
  if err != nil {
    return err
  }

  q := New(tx)
  err = fn(q)
  if err != nil {
    if rbErr := tx.Rollback(); rbErr != nil {
      return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
    }
    return err
  }

  return tx.Commit()
}

type TransferTxParams struct {
  FromAccountId int64 `json:"from_account_id"`
  ToAccountId int64 `json:"to_account_id"`
  Amount int64 `json:"amount"`
}

type TransferTxResult struct {
  Transfer Transfer `json:"Transfer"`
  FromAccount Account `json:"from_account"`
  ToAccount Account `json:"to_account"`
  FromEntry Entry `json:"from_Entry"`
  toEntry Entry `json:"to_Entry"`
}

func (store *Store) Transfertx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {

}
