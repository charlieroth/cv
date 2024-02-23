package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charlieroth/cv/internal/cv"
	"github.com/charlieroth/cv/internal/md"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		Usage()
		os.Exit(1)
		return
	}

	if args[1] == "gen" && args[2] == "md" {
		cv, err := cv.Parse("/Users/charlie/github.com/charlieroth/cv/cv.json")
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		mdString, err := md.Generate(&cv)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		os.WriteFile("cv.md", []byte(mdString), 0666)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
		return
	}

	Usage()
	os.Exit(1)
}

func Usage() {
	fmt.Println("Usage: cv gen md")
}

func GenerateMarkdown() {
	fmt.Println("Markdown generation not implemented")
}
