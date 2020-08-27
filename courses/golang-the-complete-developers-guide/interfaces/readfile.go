package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := os.Args[1]

	file, err := os.Open(fn)

	if err != nil {
		fmt.Println("Could not open the file", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
