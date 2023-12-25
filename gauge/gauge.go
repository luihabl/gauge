package gauge

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/google/go-github/v57/github"
)

var lock = sync.RWMutex{}

func SortMap(m map[string]int) []string {

	pairs := make([][2]interface{}, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1].(int) > pairs[j][1].(int)
	})

	keys := make([]string, len(pairs))
	for i, p := range pairs {
		keys[i] = p[0].(string)
	}

	return keys
}

func readRepo(userName string, repoName string, client *github.Client, langs map[string]int, wg *sync.WaitGroup) {
	rlangs, _, _ := client.Repositories.ListLanguages(context.Background(), userName, repoName)

	lock.Lock()
	defer lock.Unlock()
	defer wg.Done()

	fmt.Printf("    🏷️ Reading %s/%s\n", userName, repoName)

	for k, v := range rlangs {
		langs[k] += v
	}
}

func FetchLangs(userName string, token string) map[string]int {

	client := github.NewClient(nil)

	if len(token) > 0 {
		client = client.WithAuthToken(token)
	}

	opt := &github.RepositoryListByUserOptions{Type: "owner"}
	repos, _, err := client.Repositories.ListByUser(context.Background(), userName, opt)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	langs := make(map[string]int)

	var wg sync.WaitGroup

	for _, repo := range repos {
		wg.Add(1)
		go readRepo(userName, *repo.Name, client, langs, &wg)
	}

	wg.Wait()

	return langs
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
