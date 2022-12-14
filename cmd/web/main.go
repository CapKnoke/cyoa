package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CapKnoke/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "The port to start the web app on")
	filename := flag.String("file", "gopher.json", "Path to .json file containing story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		fmt.Println("failed to open .json file")
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
