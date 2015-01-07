package main

import (
	//"flag"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/inconshreveable/axiom"
	_ "github.com/mattn/go-sqlite3"
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
				fmt.Println("must do stuff, can this take trailing arguments?")
				fmt.Println(argsWithoutFlags)
			},
		},
		{
			Name: "should",
			Action: func(c *cli.Context) {
				fmt.Println("should do stuff, can this take trailing arguments?")
			},
		},
		{
			Name: "did",
			Action: func(c *cli.Context) {
				fmt.Println("should do stuff, can this take trailing arguments?")
			},
		},
	}

	//axiom.WrapApp(app, axiom.NewMousetrap(), axiom.NewLogged())
	axiom.WrapApp(app, axiom.NewLogged())

	app.Run(os.Args)
	/*
		datePtr := flag.String("date", today.Format(time.RFC3339), "override today, defaults to… today.")

		flag.Parse()

		argsWithoutFlags := flag.Args()

		fmt.Println("date: ", *datePtr)
		fmt.Println("tail: ", argsWithoutFlags)
	*/
}
