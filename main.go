package main

import (
	"fmt"
	"os"

	"github.com/estratocloud/manifesting/manifesting/cli"
)

func main() {
	err := cli.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "🚨", "\u001B[31mERROR:", err, "\033[0m")
		os.Exit(1)
	}
}
