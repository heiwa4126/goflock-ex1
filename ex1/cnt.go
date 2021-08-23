package ex1

// flock()を使ってるのでUNIXでしか動きません

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"unsafe"

	"github.com/heiwa4126/goflock-ex1/util"
)

const count_file = "/tmp/goflock-ex1-count1"

type Counter1 struct {
	countFile string
}

func New() *Counter1 {
	return &Counter1{countFile: count_file}
}

func (counter *Counter1) InitCounter() (cnt uint64, err error) {
	f, err := os.Create(counter.countFile)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.WriteString("0")
	return
}

func (counter *Counter1) incCounter(lock bool) (cnt uint64, err error) {
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

func (counter *Counter1) IncCounter10000(lock bool) (cnt uint64, err error) {
	for i := 0; i < 10000; i++ {
		cnt, err = counter.incCounter(lock)
		if err != nil {
			break
		}
	}
	return
}
