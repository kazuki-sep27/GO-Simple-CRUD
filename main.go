package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kazuki-sep27/simple_bank_go/api"
	db "github.com/kazuki-sep27/simple_bank_go/db/sqlc"
)

const (
	dbDriver = "mysql"
	dbSource = "admin:Qwe12345@tcp(127.0.0.1:3306)/simple_bank?parseTime=true"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("can not connect db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}