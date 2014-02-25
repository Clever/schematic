package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/heroku/schematic/schema"
	"log"
	"os"
)

var file = flag.String("file", "", "schema file")


func main() {
	log.SetFlags(0)
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}

	var s schema.Schema
	d := json.NewDecoder(f)
	if err := d.Decode(&s); err != nil {
		log.Fatal(err)
	}

	code, err := s.Generate()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(code))
}