package models

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	log "gopkg.in/inconshreveable/log15.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

var dbMap = initDb()

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
		log.Crit(msg, "models", err)
	}
}
