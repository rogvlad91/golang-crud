package main

import (
	"context"
	"golang-crud/apps/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli.RunCli(ctx)
}
