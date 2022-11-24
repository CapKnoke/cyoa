package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CapKnoke/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "Path to .json file containing story")
	entryChapter := flag.String("entry", "intro", "Chapter name to use as entry point in story. Defaults to 'intro'")
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
	reader := bufio.NewReader(os.Stdin)
	chapter := story[*entryChapter]
	runChapter(story, chapter, reader)
}

func runChapter(s cyoa.Story, c cyoa.Chapter, r *bufio.Reader) {
	opts := make(map[string]string)
	fmt.Printf("%s\n\n", c.Title)
	for _, p := range c.Paragraphs {
		fmt.Println(p)
	}
	fmt.Println("")
	if len(c.Options) == 0 {
		fmt.Print("That's all folks! press 'enter' to exit")
		_, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}
	for i, opt := range c.Options {
		fmt.Printf("-[%d]	%s\n", i+1, opt.Text)
		opts[strconv.Itoa(i+1)] = opt.Chapter
	}
	action, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	runChapter(s, s[opts[strings.TrimSpace(action)]], r)
}
