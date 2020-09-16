// +build debug

package utils

import (
	"io"
	"os"
	"strconv"
)

// IoWriter returns temporary file.
func IoWriter() (w io.Writer, err error) {
	pid := os.Getpid()
	tmp := os.TempDir() + "/"
	fname := "cuishark-" + strconv.Itoa(pid)
	fullname := tmp + fname

	w, err = os.Create(fullname)
	if err != nil {
		return nil, err
	}
	return w, err
}
