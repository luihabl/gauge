package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

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

	fmt.Println("🛜 Fetching information of user", userName)

	langs := gauge.FetchLangs(userName, token)

	sortedLangs := gauge.SortMap(langs)

	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent)

	fmt.Printf("\n --- Language weight --- \n\n")
	for _, l := range sortedLangs {
		fmt.Fprintf(w, "%s\t%s\t\n", l, gauge.ByteCountIEC(int64(langs[l])))
	}

	w.Flush()
}
