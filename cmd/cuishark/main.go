package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/n-hachi/cuishark/internal/cuishark"
	"github.com/n-hachi/cuishark/internal/utils"
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
	// Generate logger
	w, err := utils.IoWriter()
	if err != nil {
		panic(err)
	}
	log.SetOutput(w)

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
