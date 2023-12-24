package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func save(url string, wg *sync.WaitGroup, client *http.Client) error {
	defer wg.Done()

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error getting URL: %w", err)
	}
	defer resp.Body.Close()

	fileFullPath := "./lib/res/" + resp.Request.URL.Host + ".html"

	f, err := os.Create(fileFullPath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer f.Close()

	// direct stream the response body to the file
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

// TODO: error handling thru channels
func ConcurrentScrapper() {
	client := &http.Client{}
	wg := &sync.WaitGroup{}

	urls := []string{
		"https://motherfuckingwebsite.com/",
		"http://bettermotherfuckingwebsite.com",
		"https://evenbettermotherfucking.website",
		"https://fdemir.dev",
	}

	for _, url := range urls {
		wg.Add(1)
		go save(url, wg, client)
	}

	wg.Wait()

	fmt.Println("Finished!")
}
