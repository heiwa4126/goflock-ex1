package main

import (
	"log"
	"os"
	"syscall"
)

const count_file = "./var/count"

func setFileLock(f *os.File, lock bool) error {
	how := syscall.LOCK_UN
	if lock {
		how = syscall.LOCK_EX
	}
	return syscall.Flock(int(f.Fd()), how|syscall.LOCK_NB)
}

func initCounter() error {
	f, err := os.Create(count_file)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("0")
	return err
}

func incCounter10000() error {
	return nil
}

func main() {
	var err error

	if len(os.Args) == 1 || os.Args[1] == "init" {
		err = initCounter()
	} else if os.Args[1] == "inc" {
		err = incCounter10000()
	}

	if err != nil {
		log.Fatal(err)
	}
}
