package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/akihisa1210/career"
	"github.com/urfave/cli/v2"
)

func main() {
	var input string

	app := &cli.App{
		Name:  "cvgen",
		Usage: "generate curriculum vitae",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "input from `FILE`",
				Destination: &input,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			f, err := os.Open(input)
			if err != nil {
				return cli.Exit(err, 1)
			}
			defer f.Close()

			data, err := ioutil.ReadAll(f)
			if err != nil {
				return cli.Exit(err, 1)
			}

			fmt.Println(string(data)) // debug

			r := bytes.NewReader(data)

			cr, err := career.Parse(r)
			if err != nil {
				return cli.Exit(err, 1)
			}

			html, err := career.Generate(cr)
			if err != nil {
				return cli.Exit(err, 1)
			}

			fmt.Println(html) // debug

			out, err := os.Create("../../index.html")
			if err != nil {
				return cli.Exit(err, 1)
			}
			defer out.Close()
			out.Write([]byte(html))

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
