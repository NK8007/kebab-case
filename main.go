package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: kebab-case <text>")
		os.Exit(1)
	}
	fmt.Println(ToKebabCase(strings.Join(os.Args[1:], " ")))
}
