package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/02_call_apps/internal/ftests"
)

const maxExecutionTime = 5 * time.Minute

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), maxExecutionTime)
	defer cancel()

	dir := getDir()

	coordinator := &ftests.Coordinator{}

	err := coordinator.Run(ctx, dir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(-1)
	}
}

func getDir() string {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Please supply a base directory")
		os.Exit(-1)
	}

	return args[0]
}
