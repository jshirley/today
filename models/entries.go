package models

import (
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

type Entry struct {
	Id        int64  `db:"id"`
	Created   string `db:"created_at"`
	Category  string
	Title     string
	Note      string
	Completed bool
	Sentiment int64
}

func EntriesForDate(date time.Time) []Entry {
	var entries []Entry

	_, err := dbMap.Select(&entries, "SELECT * FROM entries WHERE DATE(created_at, 'localtime') = ?", date.Format("2006-01-02"))
	checkErr(err, "Unable to select entries")

	log.Debug("entry count for date", "models", len(entries), "date", date)

	return entries
}

func EntriesForToday() []Entry {
	return EntriesForDate(time.Now())
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

func AddEntryForToday(category string, entry string, note string) {
	entity := CreateEntry(category, entry, note, false)

	err := dbMap.Insert(&entity)
	checkErr(err, "Unable to save entry for today")

	log.Debug("Saved entity into the database", "entries")
}

func AddCompletedEntryForToday(category string, entry string, note string) {
	entity := CreateEntry(category, entry, note, true)

	err := dbMap.Insert(&entity)
	checkErr(err, "Unable to save entry for today")

	log.Debug("Saved entity into the database", "entries")
}
