package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/shokujinjp/markov-menu/pkg/markov"

	"github.com/shokujinjp/shokujinjp-sdk-go/shokujinjp"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	rand.Seed(time.Now().Unix())

	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return fmt.Errorf("failed to tokenizer.New: %w", err)
	}

	all, err := shokujinjp.GetMenuAllData()
	if err != nil {
		return fmt.Errorf("failed to GetMenuAllData: %w", err)
	}

	var parsed [][]string
	for _, menu := range all {
		seg := t.Wakati(menu.Name)
		block := markov.ParseMenu(seg)
		parsed = append(parsed, block...)
	}

	chain := markov.GenerateChain(parsed)
	chain = markov.TrimSystemMessages(chain)
	fmt.Println(strings.Join(chain, ""))

	return nil
}
