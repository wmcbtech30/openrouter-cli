package main

import (
	"context"
	"fmt"
	"os"

	"github.com/wmcbtech30/openrouter-cli/internal/subcommands"
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
	case "--version", "-v", "version":
		showVersion()
	case "chat":
		if err := subcommands.ChatSubCmdParser(ctx, os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		// pass the context to the subcommand
		showHelp()
		os.Exit(1)
	}
}
