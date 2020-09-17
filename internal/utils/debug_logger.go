// +build debug

package utils

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
)

// IoWriter returns temporary file.
func IoWriter() (w io.Writer, err error) {
	tmpdir := os.TempDir() + "/"
	fname := "cuishark-" + strconv.Itoa(os.Getpid())
	fullname := filepath.FromSlash(tmpdir + fname)

	w, err = os.Create(fullname)
	if err != nil {
		return nil, err
	}
	return w, err
}
