package gauge

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/google/go-github/v57/github"
)

func SortMap(m map[string]int) []string {
	// Create slice of key-value pairs
	pairs := make([][2]interface{}, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	// Sort slice based on values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1].(int) > pairs[j][1].(int)
	})

	// Extract sorted keys
	keys := make([]string, len(pairs))
	for i, p := range pairs {
		keys[i] = p[0].(string)
	}

	return keys
}

func FetchLangs(userName string, token string) map[string]int {

	client := github.NewClient(nil).WithAuthToken(token)
	opt := &github.RepositoryListByUserOptions{Type: "all"}
	repos, _, err := client.Repositories.ListByUser(context.Background(), userName, opt)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	langs := make(map[string]int)

	for _, repo := range repos {
		rlangs, _, _ := client.Repositories.ListLanguages(context.Background(), userName, *repo.Name)

		fmt.Printf("Reading %s/%s", repo.Owner, *repo.Name)

		for k, v := range rlangs {
			langs[k] += v
		}
	}

	return langs
}
