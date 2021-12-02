package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.design/x/clipboard"
)

func main() {
	fmt.Print("Enter Windows Path: ")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
		bashPath := strings.Replace(input, `\`, `/`, -1)
		clipboard.Write(clipboard.FmtText, []byte(bashPath))
	}

}
