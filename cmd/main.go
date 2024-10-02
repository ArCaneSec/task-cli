package main

import (
	"log"

	"github.com/ArCaneSec/task-cli/internal/options"
)

func main() {
	if err := options.ParseFlags(); err != nil {
		log.Fatalln("[!] Error while parsing flags: %w", err)
	}
}
