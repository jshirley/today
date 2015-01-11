package commands

import (
	"fmt"
	"github.com/jshirley/today/models"
	"github.com/russross/blackfriday"
)

func DisplayReview() {
	entries := models.EntriesForToday()
	fmt.Println("Got entries?", entries)

	notes := models.NotesForToday()
	for _, note := range notes {
		fmt.Println(note.Title)
		fmt.Println(string(blackfriday.MarkdownBasic([]byte(note.Note))))
	}

}
