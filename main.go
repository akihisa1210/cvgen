package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Activity string

type Project struct {
	Period     string
	Team       string
	Role       string
	Technology string
	Activities []Activity
}

type Campany struct {
	Name     string
	Summary  string
	Projects []Project
}

type Career struct {
	Campanys []Campany
}

func main() {
	f, err := os.Open("./career.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	career := Career{}
	err = yaml.Unmarshal([]byte(data), &career)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Valid!")
}
