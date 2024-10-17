package main

import (
	"fmt"
	"os"

	qrcode "github.com/hashmapsdata2value/go-tinyqr"
)

func main() {
	err := qrcode.Print(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
