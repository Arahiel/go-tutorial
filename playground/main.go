package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	for _, filename := range args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, file)
		fmt.Println()
	}
}
