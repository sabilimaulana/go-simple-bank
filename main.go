package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sabilimaulana/go-simple-bank/api"
	db "github.com/sabilimaulana/go-simple-bank/db/sqlc"
	"github.com/sabilimaulana/go-simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
