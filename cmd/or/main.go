package main

import (
	"context"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		os.Exit(1)
	}

	// share context
	ctx := context.Background()

	switch os.Args[1] {
	case "--help", "-h", "help":
		showHelp()
		os.Exit(0)
	case "--version", "-v", "version":
		showVersion()
		os.Exit(0)
	default:
		// pass the context to the subcommand
		showHelp()
		os.Exit(1)
	}
}
