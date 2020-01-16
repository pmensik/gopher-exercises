package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story %s.\n", *file)
	jsonFile, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	var story cyoa.Story
	if err := json.Unmarshal(jsonFile, &story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
}
