package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

const count_file = "/tmp/goflock-ex1-count"

var (
	// Version = $(git tag --sort=-v:refname | grep '^v' | head -1 | sed 's/^v//')
	Version string = "9.9.9"
	// Revision = $(git rev-parse --short HEAD)
	Revision string = "zzzzzzz"
)

func setFileLock(f *os.File, lock bool) error {
	how := syscall.LOCK_UN
	if lock {
		how = syscall.LOCK_EX
	}
	return syscall.Flock(int(f.Fd()), how)
}

func initCounter() (cnt uint64, err error) {
	f, err := os.Create(count_file)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.WriteString("0")
	return
}

func incCounter(lock bool) (cnt uint64, err error) {
	f, err := os.OpenFile(count_file, os.O_RDWR, 0664)
	if err != nil {
		return
	}
	defer f.Close()

	if lock {
		err = setFileLock(f, true)
		if err != nil {
			return
		}
		defer func() {
			if cerr := setFileLock(f, false); err == nil {
				err = cerr
			}
		}()
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	cnt, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 64)
	if err != nil {
		return
	}
	cnt++

	err = f.Truncate(0)
	if err != nil {
		return
	}
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	_, err = f.WriteString(strconv.FormatUint(cnt, 10))

	return
}

func incCounter10000(lock bool) (cnt uint64, err error) {
	for i := 0; i < 10000; i++ {
		cnt, err = incCounter(lock)
		if err != nil {
			break
		}
	}
	return
}

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
		cnt, err = incCounter10000(false)
	case "flockinc":
		cnt, err = incCounter10000(true)
	case "init":
		cnt, err = initCounter()
	default:
		help()
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", cnt)
}
