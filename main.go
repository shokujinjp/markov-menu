package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/shokujinjp/markov-menu/pkg/markov"

	"github.com/shokujinjp/shokujinjp-sdk-go/shokujinjp"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	all, err := shokujinjp.GetMenuAllData()
	if err != nil {
		return fmt.Errorf("failed to GetMenuAllData: %w", err)
	}

	parsed, err := markov.Parse(all)
	if err != nil {
		return fmt.Errorf("failed to parse: %w", err)
	}

	chain := markov.GenerateChain(parsed)
	chain = markov.TrimSystemMessages(chain)
	fmt.Println(strings.Join(chain, ""))

	return nil
}
