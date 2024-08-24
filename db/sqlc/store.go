package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
// type Store interface {
// 	Queries
// 	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
// }

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// TransferTx perform a money transfer from one account to the other.
// It creates a transfer record, add account entries, and update accounts'
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		//TODO: update accounts' balance

		return nil
	})

	return result, err

}

// NewStore creates a new SQLStore instance
func NewStore(conn *pgxpool.Pool) *SQLStore {
	queries := New(conn) // Khởi tạo Queries bằng conn (hoặc db tương đương)
	return &SQLStore{
		connPool: conn,
		Queries:  queries,
	}
}
