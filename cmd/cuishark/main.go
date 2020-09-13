package main

import (
	"context"
	"fmt"
	"os"

	"github.com/n-hachi/cuishark/internal/cuishark"
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

	// Get pcap file path
	path := os.Args[1]
	cs, err := cuishark.New(path)
	if err != nil {
		panic(err)
	}
	defer cs.End()

	if err = cs.Run(ctx); err != nil {
		return 1
	}
	return 0
}
