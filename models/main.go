package models

import (
	"database/sql"
	"path/filepath"
	//"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

var dbMap = initDb()

type Entry struct {
	Id        int64  `db:"id"`
	Created   string `db:"created_at"`
	Category  string
	Title     string
	Note      string
	Completed bool
	Sentiment int64
}

func Hello() {
	checkErr(nil, "Hello")
}

func NewDatabase() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	err := dbmap.TruncateTables()
	checkErr(err, "Truncate Tables failed")
}

func EntriesForDate(date string) []Entry {
	var entries []Entry

	log.Println("Fetching entries for ", date)
	_, err := dbMap.Select(&entries, "SELECT * FROM entries WHERE DATE(created_at, 'localtime') = ?", date)
	checkErr(err, "Unable to select entries")

	log.Println("Got results? ", len(entries))

	return entries
}

func EntriesForToday() []Entry {
	today := time.Now()
	return EntriesForDate(today.Format("2006-01-02"))
}

func AddEntryForToday(category string, entry string) {
	entity := Entry{
		Created:   time.Now().Format(time.RFC3339Nano),
		Category:  category,
		Title:     entry,
		Note:      "",
		Completed: false,
		Sentiment: 0,
	}

	err := dbMap.Insert(&entity)
	checkErr(err, "Unable to save entry for today")

	log.Println("Saved entity into the database")
}

func appRoot() string {
	root := filepath.Join(os.Getenv("HOME"), ".today")
	os.Mkdir(root, 0776)
	return root
}

func dbFile() string {
	return filepath.Join(appRoot(), "today_db.bin")
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", dbFile())
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
