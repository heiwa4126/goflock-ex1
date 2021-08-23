package ex3

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"

	"github.com/heiwa4126/goflock-ex1/util"
)

const count_file = "/tmp/goflock-ex1-count3"

type Counter3 struct {
	countFile string
}

func New() *Counter3 {
	return &Counter3{countFile: count_file}
}

func (counter *Counter3) InitCounter() (cnt uint64, err error) {
	f, err := os.Create(counter.countFile)
	if err != nil {
		return
	}
	defer f.Close()
	err = binary.Write(f, binary.LittleEndian, cnt)
	return
}

func (counter *Counter3) incCounter(lock bool) (cnt uint64, err error) {
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
	err = binary.Read(bufio.NewReader(f), binary.LittleEndian, &cnt)
	if err != nil {
		return
	}
	cnt++

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	bw := bufio.NewWriter(f)
	defer func() {
		if cerr := bw.Flush(); err == nil {
			err = cerr
		}
	}()

	err = binary.Write(bw, binary.LittleEndian, cnt)
	return
}

func (counter *Counter3) IncCounter10000(lock bool) (cnt uint64, err error) {
	for i := 0; i < 10000; i++ {
		cnt, err = counter.incCounter(lock)
		if err != nil {
			break
		}
	}
	return
}
