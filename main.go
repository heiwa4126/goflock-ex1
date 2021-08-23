package main

import (
	"fmt"
	"log"
	"os"

	"github.com/heiwa4126/goflock-ex1/ex1"
	"github.com/heiwa4126/goflock-ex1/ex2"
	"github.com/heiwa4126/goflock-ex1/ex3"
	"github.com/heiwa4126/goflock-ex1/ex4"
)

var (
	// Version = $(git tag --sort=-v:refname | grep '^v' | head -1 | sed 's/^v//')
	Version string = "9.9.9"
	// Revision = $(git rev-parse --short HEAD)
	Revision string = "zzzzzzz"
)

func help() {
	// help or version
	fmt.Printf("goflock-ex1 %s (%s)\n", Version, Revision)
	os.Exit(2)
}

func main() {
	var err error
	var cnt uint64

	if len(os.Args) < 2 {
		help()
	}

	switch os.Args[1] {
	case "inc":
		cnt, err = ex1.IncCounter10000(false)
	case "flockinc":
		cnt, err = ex1.IncCounter10000(true)
	case "init":
		cnt, err = ex1.InitCounter()
		// ---
	case "inc2":
		cnt, err = ex2.IncCounter10000(false)
	case "flockinc2":
		cnt, err = ex2.IncCounter10000(true)
	case "init2":
		cnt, err = ex2.InitCounter()
		// ---
	case "inc3":
		cnt, err = ex3.IncCounter10000(false)
	case "flockinc3":
		cnt, err = ex3.IncCounter10000(true)
	case "init3":
		cnt, err = ex3.InitCounter()
		// ---
	case "inc4":
		cnt, err = ex4.IncCounter10000(false)
	case "flockinc4":
		cnt, err = ex4.IncCounter10000(true)
	case "init4":
		cnt, err = ex4.InitCounter()
		// ---
	default:
		help()
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", cnt)
}
