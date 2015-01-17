package commands

import (
	"fmt"
	"github.com/jshirley/today/models"
	"github.com/russross/blackfriday"
	"github.com/wsxiaoys/terminal"
	//"github.com/wsxiaoys/terminal/color"
	"time"
	//log "gopkg.in/inconshreveable/log15.v2"
)

func displayPlanningHints() {
	today := time.Now()
	today_string := today.Format(time.RFC1123)

	yesterday := today.AddDate(0, 0, -1)
	entries := models.EntriesForDate(yesterday)

	terminal.Stdout.Color("w").
		Print("It is ", today_string, " and it's time to plan!").Nl()

	if len(entries) == 0 {
		terminal.Stdout.Color("r").
			Print("Yesterday was ", yesterday.Format(time.RFC1123), ".").Nl().
			Print("- Nothing was planned yesterday, nothing to review today -")
	} else {
		terminal.Stdout.Color("g").
			Print("Review what's been done and what needs to be done:").Nl()
		for _, entry := range entries {
			DisplayEntry(entry)
		}
	}
	terminal.Stdout.Nl().Reset()

	//Reset().
	//Colorf("@{kW}Hello world\n")
	//color.Print("@rHello world\n")
}

func DisplayEntry(entry models.Entry) {
	if entry.Completed == true {
		terminal.Stdout.
			Color("g").Print("[x] ")
	} else {
		terminal.Stdout.
			Color("r").Print("[ ] ")
	}

	terminal.Stdout.
		Color("w").Print(entry.Title).Nl().
		Print("    ", string(blackfriday.MarkdownBasic([]byte(entry.Note)))).Nl()
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

	if len(entries) == 0 {
		displayPlanningHints()
	}

	notes := models.NotesForToday()
	if len(notes) > 0 {
		for _, note := range notes {
			DisplayNote(note)
		}
	}
}
