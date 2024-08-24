package sqlc

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hieupc05/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	fmt.Printf("key: %v", config.DBSource)
	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = New(connPool)

	os.Exit(m.Run())
}
