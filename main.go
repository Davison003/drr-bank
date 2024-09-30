package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/Davison003/drr-bank/api"
	db "github.com/Davison003/drr-bank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:1805@localhost:5432/drr_bank?sslmode=disable"
	serverAddr = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

}