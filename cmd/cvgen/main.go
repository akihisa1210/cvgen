package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/akihisa1210/career"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: cvgen <career.yml>")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data)) // debug

	r := bytes.NewReader(data)

	c, err := career.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	html, err := career.Generate(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(html) // debug

	out, err := os.Create("../../index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	out.Write([]byte(html))
}
