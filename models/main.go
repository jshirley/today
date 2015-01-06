package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	//"time"
)

type Entry struct {
	Id        int64 `db:"id"`
	Created   int64 `db:"created_at"`
	Category  string
	Title     string
	Note      string
	Completed bool
	Sentiment int64
}

func NewDatabase() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	err := dbmap.TruncateTables()
	checkErr(err, "Truncate Tables failed")
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "/tmp/today_db.bin")
	checkErr(err, "sql.Open failed!")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTableWithName(Entry{}, "entries").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
