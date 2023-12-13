package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/solracnet/go_finance_backend/api"
	db "github.com/solracnet/go_finance_backend/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:123456@localhost:5432/go_finance?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot open database connection")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api server", err)
	}
}
