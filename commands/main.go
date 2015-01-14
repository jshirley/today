package commands

import (
	"fmt"
	"github.com/jshirley/today/models"
	"github.com/russross/blackfriday"
)

func DisplayEntry(entry models.Entry) {
	fmt.Println(entry.Title)
	fmt.Println(string(blackfriday.MarkdownBasic([]byte(entry.Note))))
}

func DisplayNote(entry models.Note) {
	fmt.Println(entry.Title)
	fmt.Println(string(blackfriday.MarkdownBasic([]byte(entry.Note))))
}

func DisplayReview() {
	entries := models.EntriesForToday()
	for _, entry := range entries {
		DisplayEntry(entry)
	}

	notes := models.NotesForToday()
	for _, note := range notes {
		DisplayNote(note)
	}

}
