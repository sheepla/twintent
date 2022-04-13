package main

import (
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/toqueteos/webbrowser"
)

func main() {
	text := strings.Join(os.Args[1:], " ")
	err := webbrowser.Open(newURL(text))
	if err != nil {
		log.Fatal(err)
	}
}

func newURL(text string) string {
	u := &url.URL{
		Scheme: "https",
		Host:   "twitter.com",
		Path:   "intent/tweet",
	}
	q := u.Query()
	q.Set("text", text)
	u.RawQuery = q.Encode()
	return u.String()
}
