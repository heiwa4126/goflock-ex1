package ex4

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/gofrs/flock"
)

const count_file = "/tmp/goflock-ex1-count4"
const lock_file = "/tmp/goflock-ex1-count4.lock"

type Counter4 struct {
	countFile string
	lockFile  string
}

func New() *Counter4 {
	return &Counter4{countFile: count_file, lockFile: lock_file}
}

func (counter *Counter4) InitCounter() (cnt uint64, err error) {
	f, err := os.Create(counter.countFile)
	if err != nil {
		return
	}
	defer f.Close()
	err = binary.Write(f, binary.LittleEndian, cnt)
	return
}

func (counter *Counter4) incCounter(lock bool) (cnt uint64, err error) {
	f, err := os.OpenFile(counter.countFile, os.O_RDWR, 0664)
	if err != nil {
		return
	}
	defer f.Close()

	if lock {
		fileLock := flock.New(counter.lockFile)
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

func (counter *Counter4) IncCounter10000(lock bool) (cnt uint64, err error) {
	for i := 0; i < 10000; i++ {
		cnt, err = counter.incCounter(lock)
		if err != nil {
			break
		}
	}
	return
}
