package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/trancongthanh1992/samplebank/api"
	db "github.com/trancongthanh1992/samplebank/db/sqlc"
	"github.com/trancongthanh1992/samplebank/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
