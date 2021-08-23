package ex2

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/heiwa4126/goflock-ex1/util"
)

const count_file = "/tmp/goflock-ex1-count2"

type Counter2 struct {
	countFile string
}

func New() *Counter2 {
	return &Counter2{countFile: count_file}
}

func (counter *Counter2) InitCounter() (cnt uint64, err error) {
	f, err := os.Create(counter.countFile)
	if err != nil {
		return
	}
	defer f.Close()
	err = binary.Write(f, binary.LittleEndian, cnt)
	return
}

func (counter *Counter2) incCounter(lock bool) (cnt uint64, err error) {
	f, err := os.OpenFile(counter.countFile, os.O_RDWR, 0664)
	if err != nil {
		return
	}
	defer f.Close()

	if lock {
		err = util.SetFileLock(f, true)
		if err != nil {
			return
		}
		defer func() {
			if cerr := util.SetFileLock(f, false); err == nil {
				err = cerr
			}
		}()
	}
	err = binary.Read(f, binary.LittleEndian, &cnt)
	if err != nil {
		return
	}
	cnt++

	// err = f.Truncate(0)
	// if err != nil {
	// 	return
	// }
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	err = binary.Write(f, binary.LittleEndian, cnt)
	return
}

func (counter *Counter2) IncCounter10000(lock bool) (cnt uint64, err error) {
	for i := 0; i < 10000; i++ {
		cnt, err = counter.incCounter(lock)
		if err != nil {
			break
		}
	}
	return
}
