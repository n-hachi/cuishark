// +build debug

package utils

import (
	"io"
	"io/ioutil"
)

// IoWriter returns temporary file.
func IoWriter() (w io.Writer, err error) {
	w, err = ioutil.TempFile("", "cuishark_")
	return w, err
}
