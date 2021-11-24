package main

import (
	"fmt"
	"strings"

	"golang.design/x/clipboard"
)

func main() {
	fmt.Print("Enter Windows Path: ")
	var input string
	fmt.Scanln(&input)
	bashPath := strings.Replace(input, `\`, `/`, -1)
	clipboard.Write(clipboard.FmtText, []byte(bashPath))
}
