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
	var (
		input  string
		output string
	)

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
			&cli.StringFlag{
				Name:        "input",
				Aliases:     []string{"i"},
				Usage:       "input from `FILE`",
				Destination: &input,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "output to `FILE`",
				Destination: &output,
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

			r := bytes.NewReader(data)

			cr, err := career.Parse(r)
			if err != nil {
				return cli.Exit(err, 1)
			}

			html, err := career.Generate(cr)
			if err != nil {
				return cli.Exit(err, 1)
			}

			if output == "" {
				fmt.Println(html)
				return nil
			}

			out, err := os.Create(output)
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
