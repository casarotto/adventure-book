package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	adventurebook "github.com/casarotto/adventure-book"
)

func main() {
	port := flag.Int("port", 3000, "the port to start CYOA Web Application")
	file := flag.String("file", "gopher.json", "JSON story file")
	flag.Parse()
	fmt.Println(*file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := adventurebook.JsonStory(f)
	if err != nil {
		panic(err)
	}

	handler := adventurebook.NewHandler(story)
	fmt.Println("Your adventure awaits, open your browser and visit localhost:3000 to begin")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
