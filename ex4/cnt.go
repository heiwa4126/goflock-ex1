package ex4

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/gofrs/flock"
)

const count_file = "/tmp/goflock-ex1-count4"
const lock_file = "/tmp/goflock-ex1-count4.lock"

func InitCounter() (cnt uint64, err error) {
	f, err := os.Create(count_file)
	if err != nil {
		return
	}
	defer f.Close()
	err = binary.Write(f, binary.LittleEndian, cnt)
	return
}

func incCounter(lock bool) (cnt uint64, err error) {
	f, err := os.OpenFile(count_file, os.O_RDWR, 0664)
	if err != nil {
		return
	}
	defer f.Close()

	if lock {
		fileLock := flock.New(lock_file)
		err = fileLock.Lock()
		if err != nil {
			return
		}
		defer func() {
			if cerr := fileLock.Unlock(); err == nil {
				err = cerr
			}
		}()
	}
	err = binary.Read(f, binary.LittleEndian, &cnt)
	if err != nil {
		return
	}
	cnt++

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	err = binary.Write(f, binary.LittleEndian, cnt)
	return
}

func IncCounter10000(lock bool) (cnt uint64, err error) {
	for i := 0; i < 10000; i++ {
		cnt, err = incCounter(lock)
		if err != nil {
			break
		}
	}
	return
}
