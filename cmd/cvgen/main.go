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
	f, err := os.Open("../../career.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	r := bytes.NewReader(data)

	c, err := career.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	career.Generate(c)
}
