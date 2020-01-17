package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"story"
)

func main() {
	port := flag.Int("port", 3000, "port to start the server on")
	file := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story %s.\n", *file)
	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	s, err := story.JsonStory(bytes)
	if err != nil {
		fmt.Println("Error parsing the story")
		os.Exit(1)
	}
	h := story.NewHandler(s)
	fmt.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
