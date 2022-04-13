package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/toqueteos/webbrowser"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErrArgs
	exitCodeErrWebBrowser
)

func main() {
	os.Exit(int(Main(os.Args[1:])))
}

func Main(args []string) exitCode {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Must require arguments.")
		return exitCodeErrArgs
	}
	text := strings.Join(args, " ")
	err := webbrowser.Open(newURL(text))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open URL: %s\n", err)
		return exitCodeErrWebBrowser
	}

	return exitCodeOK
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
