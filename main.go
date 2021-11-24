package main

import (
	"flag"
	"strings"

	"golang.design/x/clipboard"
)

func main() {
	inputPath := flag.String(`path`, `C:\Users`, `provide a path from windows file system to convert them to bash format`)

	flag.Parse()

	bashPath := strings.Replace(*inputPath, `\`, `/`, -1)
	clipboard.Write(clipboard.FmtText, []byte(bashPath))
}
