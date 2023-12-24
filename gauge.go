package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v57/github"
	"github.com/joho/godotenv"
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

	client := github.NewClient(nil).WithAuthToken(token)
	opt := &github.RepositoryListByUserOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByUser(context.Background(), userName, opt)

	if err != nil {
		log.Fatal(err)
		return
	}

	langs := make(map[string]int)

	for _, repo := range repos {
		rlangs, _, _ := client.Repositories.ListLanguages(context.Background(), userName, *repo.Name)

		fmt.Printf("Reading %s/%s", repo.Owner, *repo.Name)

		for k, v := range rlangs {
			langs[k] += v
		}
	}

	for k, v := range langs {
		fmt.Printf("%s: %d KB\n", k, v/1e3)
	}

}
