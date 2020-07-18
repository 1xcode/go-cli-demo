package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			firstname := ""
			midname := ""
			lastname := ""
			if c.NArg() > 0 {
				firstname = c.Args().Get(0)
				midname = c.Args().Get(1)
				lastname = c.Args().Get(2)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("Hola", firstname, midname, lastname)
			} else {
				fmt.Println("Hello", firstname, midname, lastname)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
