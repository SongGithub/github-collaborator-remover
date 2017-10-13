package main

import (
	"flag"
	"fmt"
	"os"
	// "github.com/MYOB-Technology/ops-repo-toolbox/pkg/collaborator-remover"
)

var org = flag.String("owner", "MYOB-Technology", "name of the repo owner")

func die(s string, err int) {
	fmt.Fprintf(os.Stderr, s+"\n")
	os.Exit(err)
}

func usage() {
    String := `clrt-rm is a Github tool
        Usage: %s [flags]
    `
    fmt.Fprintf(os.Stderr, String, os.Args[0])
    flag.PrintDefaults()
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		die("GITHUB_TOKEN not set...", 1)
	}

	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
        usage()
		die("Please give me some arguments, check usage", 1)
	}
}
