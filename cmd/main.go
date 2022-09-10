package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "relayer",
		Usage: "Ethereum transaction relayer",
		Action: func(*cli.Context) error {
			fmt.Println("unimplemented")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
