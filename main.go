package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/heiwa4126/goflock-ex1/ex1"
	// "github.com/heiwa4126/goflock-ex1/ex2"
	// "github.com/heiwa4126/goflock-ex1/ex3"
	// "github.com/heiwa4126/goflock-ex1/ex4"
)

type CounterAction int

const (
	COUNTER_INIT CounterAction = iota
	COUNTER_INC
	COUNTER_LINC
)
const (
	CMD_INIT = "init"
	CMD_INC  = "inc"
	CMD_LINC = "flockinc"
)

type Counter interface {
	InitCounter() (uint64, error)
	IncCounter10000(lock bool) (uint64, error)
}

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

var cmdint_re = regexp.MustCompile(`(\D+)(\d*)`)

func cmdint(cmd string) (string, int64) {
	// ex: "int64" -> "int",64
	// ex: "int" -> "int",1 (1 as default)
	r := cmdint_re.FindStringSubmatch(cmd)
	if r == nil {
		return "", 1
	}

	n, err := strconv.ParseInt(r[2], 10, 64)
	if err != nil {
		n = 1
	}

	return r[1], n
}

func parseCmdline(args []string) (*Counter, CounterAction) {
	if len(args) < 2 {
		help()
	}

	//var counter *Counter
	counter := ex1.New()
	cAct := COUNTER_INIT
	cmd, _ := cmdint(args[1])
	// cmd, exnum := cmdint(args[1])
	// switch exnum {
	// case 1:
	// default:
	// 	help()
	// }
	switch cmd {
	case CMD_INIT:
		cAct = COUNTER_INIT
	case CMD_INC:
		cAct = COUNTER_INC
	case CMD_LINC:
		cAct = COUNTER_LINC
	default:
		help()
	}

	return counter, cAct
}

func main() {
	var err error
	var cnt uint64

	counter := ex1.New()
	//counter, cAct := parseCmdline(os.Args)
	fmt.Printf("counter: %v\n", counter)

	switch cAct {
	case COUNTER_INIT:
		cnt, err = counter.InitCounter()
	case COUNTER_INC:
		cnt, err = counter.IncCounter10000(false)
	case COUNTER_LINC:
		cnt, err = counter.IncCounter10000(true)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", cnt)
}
