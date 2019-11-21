package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/cyoa/story"
	"log"
	"net/http"
	"os"
)

func main() {
	fileName := flag.String("story", "gopher.json", "Json file with the story file")
	templateFile := flag.String("template", "index.html", "Template file for rendering the story on web")
	Port := flag.Int("Port", 8080, "port number to run web app")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		os.Exit(2)
	}
	story, err := cyoa.JsonifyStory(file)
	if err != nil {
		panic(err)
	}

	tpl := cyoa.MakeTemplate(*templateFile)
	handler := cyoa.NewStoryHandler(story, tpl)
	fmt.Printf("Starting web application serve: %d", *Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *Port), handler))
}
