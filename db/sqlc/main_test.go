package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testQuery *Queries
var testDB *sql.DB

const (
	dbDriver	=	"mysql"
	dbSource	=	"admin:Qwe12345@tcp(127.0.0.1:3306)/simple_bank?parseTime=true"
)

func TestMain(m *testing.M) {
	var err error

	testDB,err	= 	sql.Open(dbDriver,dbSource)

	if err != nil {
		log.Fatal("can not connect db:",err)
	}

	testQuery = New(testDB)

	os.Exit(m.Run())
}