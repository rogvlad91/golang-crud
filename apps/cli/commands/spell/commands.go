package spell

import (
	"context"
	"fmt"
)

type SpellCommand struct {
}

func (c *SpellCommand) Validate(args []string) {
}

func (c *SpellCommand) Execute(ctx context.Context, args []string) error {
	word := args[2]
	for i := 0; i < len(word); i++ {
		fmt.Printf("%s ", string(word[i]))
	}
	fmt.Println()
	return nil
}

func (c *SpellCommand) Usage() {
	fmt.Println("spell <word>")
}
