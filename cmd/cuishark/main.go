package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cs := cuishark.New()

	err := cs.Init()
	if err != nil {
		return 1
	}
	defer cs.End()

	if err = cs.Run(ctx); cs != nil {
		return 1
	}
	return 0
}
