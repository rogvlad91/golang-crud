package main

import (
	"context"
	"homework-7/apps/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli.RunCli(ctx)
}
