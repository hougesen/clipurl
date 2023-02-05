package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	fmt.Println("masd")
}

func checkClipboard() {

}

func findUrls(text string) []string {
	if len(text) > 0 {
		re, _ := regexp.Compile(`[-a-zA-Z0-9@:%_\+.~#?&//=]{2,256}\.[a-z]{2,4}\b(\/[-a-zA-Z0-9@:%_\+.~#?&//=]*)?`)

		urls := re.FindAllString(text, -1)

		if len(urls) > 1 {
			seen_urls := map[string]bool{}

			for i := range urls {
				index := len(urls) - 1 - i
				url := urls[index]

				if _, ok := seen_urls[url]; ok {
					urls[index] = urls[len(urls)-1]
					urls[len(urls)-1] = ""
					urls = urls[:len(urls)-1]
				}

				seen_urls[url] = true
			}
		}

		return urls
	}

	return make([]string, 0)
}

func updateSavedUrls(urls []string) {
	f, err := os.OpenFile("clipboard_urls.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()

	timestamp := time.Now().Unix()

	for _, url := range urls {
		f.WriteString(fmt.Sprintf("%d %s\n", timestamp, url))
	}

}
