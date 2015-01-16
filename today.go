package main

import (
	//"flag"
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/inconshreveable/axiom"
	_ "github.com/mattn/go-sqlite3"
	//log "gopkg.in/inconshreveable/log15.v2"
	"os"
	"strings"
	//"time"
	"github.com/jshirley/today/commands"
	"github.com/jshirley/today/models"
	"github.com/jshirley/today/server"
)

func main() {
	//today := time.Now()

	//models.NewDatabase()

	app := cli.NewApp()
	app.Name = "today"
	app.Usage = "proactive and reactive daily achievements"
	app.Author = "Jay Shirley"
	app.Email = "help@mustshouldwant.today"
	app.Version = "0.1.0"

	messageFlag := []cli.Flag{
		/* TODO: What I'd love here is something that matches git's -m.
		 * If `today did -m` then open $EDITOR, but if a string is set use that
		 */
		cli.BoolFlag{
			Name:  "message, m",
			Usage: "additional message to attach, markdown supported.",
		},
	}

	app.Action = func(c *cli.Context) {
		commands.DisplayReview()
	}

	app.Commands = []cli.Command{
		{
			Name: "review",
			Action: func(c *cli.Context) {
				commands.DisplayReview()
			},
		},
		{
			Name: "must",
			Action: func(c *cli.Context) {
				argsWithoutFlags := c.Args()
				if len(argsWithoutFlags) > 0 {
					models.AddEntryForToday("must", strings.Join(argsWithoutFlags, " "), "")
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
			Name:  "do",
			Flags: messageFlag,
			Action: func(c *cli.Context) {
				entry, message := extractEntryAndMessage(c)
				models.AddEntryForToday("", entry, message)
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
		{
			Name: "server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bind, b",
					Usage: "bind address, defaults to :8100",
					Value: "0.0.0.0:8100",
				},
			},
			Action: func(c *cli.Context) {
				server.RunTodayServer(c.String("bind"))
			},
		},
	}

	//axiom.WrapApp(app, axiom.NewMousetrap(), axiom.NewLogged())
	axiom.WrapApp(app, axiom.NewLogged())

	//TODO: add axiom.VersionCommand()
	app.Run(os.Args)
}

func extractEntryAndMessage(c *cli.Context) (string, string) {
	argsWithoutFlags := c.Args()
	if len(argsWithoutFlags) == 0 {
		abort(c, errors.New("must specify an message, such as `today do take out the trash`"))
	}
	entry := strings.Join(argsWithoutFlags, " ")

	message := ""
	if c.Bool("message") == true {
		message = models.MessageFromEditor()
	}
	return entry, message
}

func abort(c *cli.Context, err error) {
	cli.ShowCommandHelp(c, c.Command.Name)
	fmt.Println("Command error:", err)
	os.Exit(1)
}
