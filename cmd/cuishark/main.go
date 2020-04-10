package main

import (
	"fmt"
	"os"

	"github.com/n-hachi/go-cuishark/internal/cuishark"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error %s\n", err)
			os.Exit(1)
		}
	}()

	os.Exit(_main())
}

func _main() int {
	cs := cuishark.New()

	defer cuishark.End()
	err := cs.Init()
	if err != nil {
		return 1
	}
	return 0
}
