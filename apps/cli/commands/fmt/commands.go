package fmt

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type FmtCommand struct {
}

func (c *FmtCommand) Validate(args []string) {
}

func (c *FmtCommand) Execute(ctx context.Context, args []string) error {
	exceptionsWordMap := make(map[string]int)
	exceptionsWordMap["Mr"] = 1
	exceptionsWordMap["Russia"] = 2
	exceptionsWordMap["EMINEM"] = 3
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filePath := os.Args[2]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			text = text + "\n\n"
			continue
		}
		words := strings.Fields(line)
		sentence := "\t" + words[0] + " "
		for _, word := range words[1:] {
			if unicode.IsUpper(rune(word[0])) && !checkException(word, exceptionsWordMap) {
				text += strings.Trim(sentence, " ") + ". "
				sentence = word + " "
			} else {
				sentence += word + " "
			}
		}
		if sentence != "" {
			text += strings.Trim(sentence, " ") + ". "
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(text)
	return nil
}

func (c *FmtCommand) Usage() {
	fmt.Println("fmt <filepath>")
}

func checkException(word string, exceptionMap map[string]int) bool {
	_, ok := exceptionMap[word]
	return ok
}
