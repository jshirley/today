package models

import (
	log "gopkg.in/inconshreveable/log15.v2"
	"time"
)

type Note struct {
	Id        int64  `db:"id"`
	Created   string `db:"created_at"`
	Title     string
	Note      string
	Sentiment int64
}

func NotesForDate(date string) []Note {
	var notes []Note

	_, err := dbMap.Select(&notes, "SELECT * FROM notes WHERE DATE(created_at, 'localtime') = ?", date)
	checkErr(err, "Unable to select notes")

	log.Debug("note count for date", "models", len(notes), "date", date)

	return notes
}

func NotesForToday() []Note {
	today := time.Now()
	return NotesForDate(today.Format("2006-01-02"))
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

	log.Debug("Saved entity into the database", "models")
}
