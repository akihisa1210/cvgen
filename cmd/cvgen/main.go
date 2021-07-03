package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akihisa1210/career"
	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:    "cvgen",
		Version: "v0.0.1",
		Usage:   "generate curriculum vitae",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "markdown",
				Aliases: []string{"m"},
				Usage:   "output as markdown",
			},
		},
		Action: func(c *cli.Context) error {
			filename := c.Args().First()
			if filename == "" {
				cli.ShowAppHelp(c)
			}

			f, err := os.Open(filename)
			if err != nil {
				return cli.Exit(err, 1)
			}
			defer f.Close()

			cr, err := career.Parse(f)
			if err != nil {
				return cli.Exit(err, 1)
			}

			var op string

			if c.Bool("markdown") {
				op, err = career.MarkDownGenerate(cr)
			} else {
				op, err = career.HTMLGenerate(cr)
			}
			if err != nil {
				return cli.Exit(err, 1)
			}

			fmt.Println(op)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
