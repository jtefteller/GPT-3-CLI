package main

import (
	"log"
	"os"

	"github.com/jtefteller/copilot_cli/app/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
