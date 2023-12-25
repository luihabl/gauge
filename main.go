package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/luihabl/gauge/gauge"
)

func main() {
	_ = godotenv.Load()

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("No arguments provided")
		os.Exit(1)
	}

	userName := args[0]
	token := os.Getenv("GITHUB_TOKEN")

	if len(userName) < 3 {
		log.Fatal("Username name too short")
		os.Exit(1)
	}

	langs := gauge.FetchLangs(userName, token)

	sortedLangs := gauge.SortMap(langs)

	for _, l := range sortedLangs {
		fmt.Printf("%s: %d\n", l, langs[l])
	}
}
