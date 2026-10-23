package subcommands

import (
	"context"
)

func ChatSubCmdParser(ctx context.Context, args []string) error {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	return nil
}
