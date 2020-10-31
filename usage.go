package main

import (
	"fmt"
	"os"
)

func usage() string {
	return fmt.Sprintf(`usage:

	%s

env vars:

	PORT	port to serve at
	DIR	directory to serve
`, os.Args[0])
}
