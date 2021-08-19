package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"unsafe"
)

const count_file = "./var/count"

// func setFileLock(f *os.File, lock bool) error {
// 	how := syscall.LOCK_UN
// 	if lock {
// 		how = syscall.LOCK_EX
// 	}
// 	return syscall.Flock(int(f.Fd()), how|syscall.LOCK_NB)
// }

func initCounter() (cnt uint64, err error) {
	f, err := os.Create(count_file)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.WriteString("0")
	return
}

func incCounter() (uint64, error) {
	f, err := os.OpenFile(count_file, os.O_RDWR, 0664)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return 0, err
	}

	cnt, err := strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 64)
	if err != nil {
		return 0, err
	}
	cnt++

	err = f.Truncate(0)
	if err != nil {
		return 0, err
	}
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return 0, err
	}

	_, err = f.WriteString(strconv.FormatUint(cnt, 10))
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func incCounter10000() (cnt uint64, err error) {
	for i := 0; i < 10000; i++ {
		cnt, err = incCounter()
		if err != nil {
			break
		}
	}
	return
}

func main() {
	var err error
	var cnt uint64

	if len(os.Args) == 1 || os.Args[1] == "inc" {
		cnt, err = incCounter10000()
		// } else if os.Args[1] == "flock_inc" {

	} else {
		// default action
		cnt, err = initCounter()
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", cnt)
}
