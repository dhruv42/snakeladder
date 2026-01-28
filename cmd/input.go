package cmd

import (
	"bufio"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func Takeinput() string {
	scanner.Scan()
	return scanner.Text()
}
