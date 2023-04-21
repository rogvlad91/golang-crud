package cli

import "context"

type Command interface {
	Validate(args []string)
	Execute(ctx context.Context, args []string) error
	Usage()
}
