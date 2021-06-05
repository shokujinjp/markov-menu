package markov

import (
	"math/rand"
	"strings"
)

// flag messages
const (
	MessageBegin = "__BEGIN__"
	MessageEnd   = "__END__"
)

// ParseMenu parse strings
func ParseMenu(wakatied []string) [][]string {
	var result [][]string
	if len(wakatied) < 3 {
		return nil
	}

	head := []string{MessageBegin, wakatied[0], wakatied[1]}
	result = append(result, head)

	for i := 1; i < len(wakatied)-2; i++ {
		block := []string{wakatied[i], wakatied[i+1], wakatied[i+2]}
		result = append(result, block)
	}

	tail := []string{wakatied[len(wakatied)-2], wakatied[len(wakatied)-1], MessageEnd}
	result = append(result, tail)

	return result
}

func findBlock(blocks [][]string, target string) [][]string {
	var result [][]string

	for _, b := range blocks {
		if b[0] == target {
			result = append(result, b)
		}
	}

	return result
}

func connectBlock(blocks [][]string, prevResult []string) []string {
	i := 0

	for _, word := range blocks[rand.Intn(len(blocks))] {
		if i != 0 {
			prevResult = append(prevResult, word)
		}
		i++
	}

	return prevResult
}

// GenerateChain generate chain using markov
func GenerateChain(parsed [][]string) []string {
	beginBlocks := findBlock(parsed, MessageBegin)
	result := connectBlock(beginBlocks, []string{})

	count := 0
	for result[len(result)-1] != MessageEnd {
		block := findBlock(parsed, result[len(result)-1])
		if len(block) == 0 {
			break
		}
		result = connectBlock(block, result)

		count++
		if count >= 150 {
			// stop infinite loop
			break
		}
	}

	return result
}

// TrimSystemMessages trim system messages
func TrimSystemMessages(chain []string) []string {
	var result []string

	for _, c := range chain {
		if !strings.EqualFold(c, MessageBegin) && !strings.EqualFold(c, MessageEnd) {
			result = append(result, c)
		}
	}

	return result
}
