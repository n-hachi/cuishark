// +build !debug

package utils

import (
	"io"
	"io/ioutil"
)

// IoWriter returns Discard io.Writer.
func IoWriter() (w io.Writer, err error) {
	w = ioutil.Discard
	return w, nil
}
