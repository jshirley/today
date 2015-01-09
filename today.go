package main

import (
	//"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/inconshreveable/axiom"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	//"log"
	"os"
	//"time"
	"github.com/jshirley/today/models"
)

func main() {
	//today := time.Now()

	//models.NewDatabase()

	app := cli.NewApp()
	app.Name = "today"
	app.Usage = "proactive and reactive daily achievements"

	messageFlag := []cli.Flag{
		/* TODO: What I'd love here is something that matches git's -m.
		 * If `today did -m` then open $EDITOR, but if a string is set use that
		 */
		cli.BoolFlag{
			Name:  "message, m",
			Usage: "additional message to attach, markdown supported.",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "review",
			Action: func(c *cli.Context) {
				// May be a date?
				argsWithoutFlags := c.Args()
				entries := models.EntriesForToday()
				fmt.Println("Got entries?", entries)
				//entries := models.EntriesForToday()
				//fmt.Println("Got entries? %v", entries)
				fmt.Println(argsWithoutFlags)
			},
		},
		{
			Name: "must",
			Action: func(c *cli.Context) {
				argsWithoutFlags := c.Args()
				if len(argsWithoutFlags) > 0 {
					models.AddEntryForToday("must", strings.Join(argsWithoutFlags, " "))
				}
			},
		},
		{
			Name: "should",
			Action: func(c *cli.Context) {
				fmt.Println("should do stuff, can this take trailing arguments?")
			},
		},
		{
			/*
			* This action, called with no arguments, should display a list of all
			* pending items for today. If no items exist, display a celebratory
			* message.
			* If additional trailing arguments are present, assume it is a new entry
			* that is completed and record it as such.
			*
			* If the -m argument is specified, that is an additional message to append
			 */
			Name:  "did",
			Flags: messageFlag,
			Action: func(c *cli.Context) {
				argsWithoutFlags := c.Args()
				if len(argsWithoutFlags) > 0 {
					message := ""
					fmt.Println("should do stuff, can this take trailing arguments?", strings.Join(argsWithoutFlags, " "))
					if c.Bool("message") == true {
						message = models.MessageFromEditor()
					}
					fmt.Println(message)
					models.AddNoteForToday(strings.Join(argsWithoutFlags, " "), message)
				}
			},
		},
	}

	//axiom.WrapApp(app, axiom.NewMousetrap(), axiom.NewLogged())
	axiom.WrapApp(app, axiom.NewLogged())

	//TODO: add axiom.VersionCommand()
	app.Run(os.Args)
}
