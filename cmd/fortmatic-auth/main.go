// cmd/fortmatic-auth/main.go
package main

import (
	"flag"
	"log"
	"os"

	"fortmatic-auth/internal/fortmatic-auth"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	input := flag.String("input", "", "Input file path")
	output := flag.String("output", "", "Output file path")
	flag.Parse()

	app := fortmatic-auth.NewApp(*verbose)
	if err := app.Run(*input, *output); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
