package models

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

type Note struct {
	Id        int64  `db:"id"`
	Created   string `db:"created_at"`
	Title     string
	Note      string
	Sentiment int64
}

func MessageFromEditor() string {
	file, err := ioutil.TempFile(os.TempDir(), "today-")
	if err != nil {
		fmt.Println("Error opening tempfile: ", err)
		return ""
	}

	fmt.Println("Execute $EDITOR on ", file.Name())
	fmt.Println(os.Getenv("EDITOR"))

	cmd := exec.Command(os.Getenv("EDITOR"), file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting editor command", err)
		return ""
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for editor", err)
		return ""
	}

	message, err := ioutil.ReadFile(file.Name())
	defer os.Remove(file.Name())

	return string(message)
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

func CreateEntry(category string, entry string, note string, completed bool) Entry {
	return Entry{
		Created:   time.Now().Format(time.RFC3339Nano),
		Category:  category,
		Title:     entry,
		Note:      note,
		Completed: completed,
		Sentiment: 0,
	}
}

func AddEntryForToday(category string, entry string) {
	entity := CreateEntry(category, entry, "", false)

	err := dbMap.Insert(&entity)
	checkErr(err, "Unable to save entry for today")

	log.Println("Saved entity into the database")
}

func AddCompletedEntryForToday(category string, entry string, note string) {
	entity := CreateEntry(category, entry, note, true)

	err := dbMap.Insert(&entity)
	checkErr(err, "Unable to save entry for today")

	log.Println("Saved entity into the database")
}

func AddNoteForToday(title string, entry string) {
	entity := Note{
		Created:   time.Now().Format(time.RFC3339Nano),
		Title:     title,
		Note:      entry,
		Sentiment: 0,
	}

	err := dbMap.Insert(&entity)
	checkErr(err, "Unable to save note for today")

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
	dbmap.AddTableWithName(Note{}, "notes").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
