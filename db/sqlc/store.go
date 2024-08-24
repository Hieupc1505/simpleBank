package sqlc

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
// type Store interface {
// 	Queries
// 	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
// }

// Store provides all functions to execute SQL queries and transactions
type Store struct {
	*Queries
	connPool *pgxpool.Pool
}

var typeKey = struct{}{}

// NewStore creates a new Store instance
func NewStore(conn *pgxpool.Pool) *Store {

	return &Store{
		connPool: conn,
		Queries:  New(conn),
	}
}
