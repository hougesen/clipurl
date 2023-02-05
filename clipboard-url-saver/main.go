package main

import (
	"fmt"
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

}
