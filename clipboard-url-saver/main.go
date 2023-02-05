package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

func main() {
	app := &cli.App{
		Name:  "clipboard-url-saver",
		Usage: "automatically saves urls found in clipboard",
		Authors: []*cli.Author{
			{
				Name:  "Mads Hougesen",
				Email: "mads@mhouge.dk",
			},
		},
		EnableBashCompletion: true,
		Suggest:              true,
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"listen", "watch"},
				Usage:   "listens to clipboard",
				Action: func(cCtx *cli.Context) error {
					listenToClipboard()
					return nil
				},
			},
			{
				Name:    "history",
				Aliases: []string{},
				Usage:   "Lists history of saved urls",
				Action: func(ctx *cli.Context) error {
					urlHistory()
					return nil
				},
			},
		},
		CommandNotFound: func(cCtx *cli.Context, command string) {
			fmt.Fprintf(cCtx.App.Writer, "Command %q not found.\n", command)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func listenToClipboard() {
	fmt.Println("Listening to clipboard")

	err := clipboard.Init()

	if err != nil {
		panic(err)
	}

	setupDirectory()

	ticker := time.NewTicker(time.Second)

	last_clipboard_check := []byte{}

	go func() {
		for range ticker.C {
			current_clipboard := getClipboard()

			comparison := bytes.Compare(last_clipboard_check, current_clipboard)

			if comparison != 0 {
				clipboard_text := string(current_clipboard)

				urls := findUrls(clipboard_text)

				if len(urls) > 0 {
					updateSavedUrls(urls)
				}

				last_clipboard_check = current_clipboard
			}
		}
	}()

	for range ticker.C {
	}
}

func getDirectoryPath() string {
	home_dir, _ := os.UserHomeDir()

	return filepath.Join(home_dir, "clipboard_urls")
}

func setupDirectory() {
	_ = os.Mkdir(getDirectoryPath(), os.ModePerm)
}

func getClipboard() []byte {
	return clipboard.Read(clipboard.FmtText)
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
	setupDirectory()

	path := filepath.Join(getDirectoryPath(), "urls.txt")

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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

func urlHistory() {
	setupDirectory()

	path := filepath.Join(getDirectoryPath(), "urls.txt")

	f, err := os.ReadFile(path)

	if err == nil && len(f) > 0 {
		fmt.Print(string(f))
	}
}
