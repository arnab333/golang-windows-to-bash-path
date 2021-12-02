package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.design/x/clipboard"
)

func main() {
	fmt.Print("Enter Windows Path: ")
	var input string
	var pathArr []string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
		bashPath := strings.Replace(input, `\`, `/`, -1)
		for i := 0; i < len(bashPath); i++ {
			pathArr = append(pathArr, string(bashPath[i]))
			if unicode.IsSpace(rune(bashPath[i])) {
				pathArr[i] = strings.Replace(string(bashPath[i]), ` `, `\ `, -1)
			}
		}
		clipboard.Write(clipboard.FmtText, []byte(strings.Join(pathArr, "")))
	}

}
