package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kazuki-sep27/simple_bank_go/util"
)

var testQuery *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Can not load config :", err)
	}	
	
	testDB,err	= 	sql.Open(config.DBDriver,config.DBSource)

	if err != nil {
		log.Fatal("can not connect db:",err)
	}

	testQuery = New(testDB)

	os.Exit(m.Run())
}