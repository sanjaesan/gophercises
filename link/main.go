package main

import (
	"flag"
	"fmt"
	"gophercises/link/parser"
	"os"
)

func main() {
	filename := flag.String("file", "ex.html", "the html file to parse")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	link, err := parser.Parse(file)
	if err != nil {
		panic(err)
	}
	for _, value := range link {
		fmt.Println("Href: ", value.Href)
		fmt.Println("Text: ", value.Text)
	}
}
