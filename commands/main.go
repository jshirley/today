package commands

import (
	"github.com/jshirley/today/models"
	"github.com/russross/blackfriday"
	"github.com/wsxiaoys/terminal"
	//"github.com/wsxiaoys/terminal/color"
	"time"
	//log "gopkg.in/inconshreveable/log15.v2"
)

func DisplayReview(timeframe string) {
	today := time.Now()
	today_string := today.Format(time.RFC1123)

	yesterday := today.AddDate(0, 0, -1)

	entries := models.EntriesForToday()
	if len(entries) == 0 {
		entries = models.EntriesForDate(yesterday)

		if len(entries) > 0 {
			terminal.Stdout.Color("g").
				Print("Yesterday", yesterday.Format(time.RFC1123), " you wanted to get this done:").Nl()
			for _, entry := range entries {
				DisplayEntry(entry)
			}
		} else {
			terminal.Stdout.Color("g").
				Print("It is ", today_string, " and it's time to plan!").Nl()
		}
	} else {
		entries := models.EntriesForDate(yesterday)

		if len(entries) == 0 {
			terminal.Stdout.Color("r").
				Print("Today is ", today_string, " and it is time to plan!").Nl().
				Color("w").
				Print("Use `must`, `should`, or `want` to track what you hope to get done today.").Nl()
		} else {
			terminal.Stdout.Color("w").
				Print("It is ", today_string, " and here's what to do:").Nl()
			for _, entry := range entries {
				DisplayEntry(entry)
			}
		}
	}
	notes := models.NotesForToday()
	if len(notes) > 0 {
		terminal.Stdout.Color("y").
			Print("Notes for ", today_string, ":").Nl()
		for _, note := range notes {
			DisplayNote(note)
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
	terminal.Stdout.Color("w").
		Print(entry.Title).Nl().
		Print(string(blackfriday.MarkdownBasic([]byte(entry.Note)))).Nl()
}
