package ex1

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"unsafe"

	"github.com/heiwa4126/goflock-ex1/util"
)

const count_file = "/tmp/goflock-ex1-count1"

func InitCounter() (cnt uint64, err error) {
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

func IncCounter10000(lock bool) (cnt uint64, err error) {
	for i := 0; i < 10000; i++ {
		cnt, err = incCounter(lock)
		if err != nil {
			break
		}
	}
	return
}
