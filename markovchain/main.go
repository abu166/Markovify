package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// MarkovChain stores prefix-suffix mappings
type MarkovChain struct {
	data map[string][]string
}

// NewMarkovChain initializes the data structure
func NewMarkovChain() *MarkovChain {
	return &MarkovChain{data: make(map[string][]string)}
}

// Build processes the input text to populate the Markov chain
func (mc *MarkovChain) Build(text string, prefixLength int) error {
	if prefixLength < 1 {
		return errors.New("prefix length must be at least 1")
	}
	words := strings.Fields(text)
	if len(words) < prefixLength+1 {
		return errors.New("not enough words to build the chain")
	}

	for i := 0; i <= len(words)-prefixLength-1; i++ {
		prefix := strings.Join(words[i:i+prefixLength], " ")
		suffix := words[i+prefixLength]
		mc.data[prefix] = append(mc.data[prefix], suffix)
	}
	return nil
}

// Generate creates random text based on the Markov chain
func (mc *MarkovChain) Generate(text, startingPrefix string, prefixLength, maxWords int) (string, error) {
	var words []string

	// Default starting prefix: Use the first `prefixLength` words from the input text
	if startingPrefix == "" {
		allWords := strings.Fields(text)
		if len(allWords) < prefixLength {
			return "", errors.New("input text does not contain enough words to form the starting prefix")
		}
		words = allWords[:prefixLength]
	} else {
		words = strings.Fields(startingPrefix)
	}

	// Check prefix length
	if len(words) != prefixLength {
		return "", fmt.Errorf("starting prefix length (%d words) does not match specified prefix length (%d words)", len(words), prefixLength)
	}

	// Generate the text
	var output []string
	output = append(output, words...)
	currentPrefix := strings.Join(words, " ")

	for len(output) < maxWords {
		suffixes, exists := mc.data[currentPrefix]
		if !exists {
			break // No further suffixes available
		}

		// Choose a random suffix
		nextWord := suffixes[rand.Intn(len(suffixes))]
		output = append(output, nextWord)

		// Update the prefix
		if prefixLength > 1 {
			words = append(words[1:], nextWord)
			currentPrefix = strings.Join(words, " ")
		} else {
			currentPrefix = nextWord
		}
	}
	return strings.Join(output, " "), nil
}

// selectRandomPrefix picks a random prefix from the Markov chain
func (mc *MarkovChain) selectRandomPrefix() (string, error) {
	if len(mc.data) == 0 {
		return "", errors.New("no data available to select a random prefix")
	}

	rand.Seed(time.Now().UnixNano())
	for prefix := range mc.data {
		return prefix, nil
	}
	return "", errors.New("unable to select a random prefix")
}

// readInput reads input from stdin and ensures there's data to process
func readInput() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil || info.Mode()&os.ModeCharDevice != 0 {
		return "", errors.New("no input text provided")
	}
	scanner := bufio.NewScanner(os.Stdin)
	var text strings.Builder
	for scanner.Scan() {
		text.WriteString(scanner.Text() + " ")
	}
	return text.String(), scanner.Err()
}

// validateInputs ensures all input arguments are valid
func validateInputs(maxWords, prefixLength int) error {
	if maxWords <= 0 {
		return errors.New("max words (-w) must be a positive integer")
	}
	if maxWords > 10000 {
		return errors.New("max words (-w) cannot exceed 10,000")
	}
	if prefixLength <= 0 {
		return errors.New("prefix length (-l) must be greater than 0")
	}
	return nil
}

// showHelp displays detailed usage instructions
func showHelp() {
	helpText := `
Markov Chain text generator.

Usage:
  markovchain [-w <N>] [-p <S>] [-l <N>]
  markovchain --help

Options:
  --help   Show this screen.
  -w N     Number of maximum words to generate. Default is 100.
  -p S     Starting prefix. If not specified, a random prefix is selected.
  -l N     Prefix length used to build the Markov chain. Default is 2.

Examples:
  Generate text with a maximum of 10 words starting with a specific prefix:
    cat input.txt | ./markovchain -w 10 -p "specific prefix"

  Generate text with a default maximum of 100 words:
    cat input.txt | ./markovchain

  Use a custom prefix length for generating text:
    cat input.txt | ./markovchain -l 3 -w 20
`
	fmt.Println(helpText)
}

// errorHandler centralizes error reporting
func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func validatePrefixLength(prefixLength int, startingPrefix string) error {
	// Check if prefixLength is valid
	if prefixLength < 1 {
		return fmt.Errorf("prefix length cannot be negative or zero")
	}
	if prefixLength > 5 {
		return fmt.Errorf("prefix length cannot be greater than 5 (given: %d)", prefixLength)
	}

	// Validate startingPrefix length if provided
	if startingPrefix != "" {
		prefixWords := strings.Fields(startingPrefix)
		if len(prefixWords) != prefixLength {
			return fmt.Errorf("starting prefix length (%d words) does not match specified prefix length (%d words)", len(prefixWords), prefixLength)
		}
	}

	return nil // No errors
}

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}()

	// Command-line arguments
	maxWords := flag.Int("w", 100, "maximum number of words to generate (default: 100, max: 10,000)")
	startingPrefix := flag.String("p", "", "starting prefix")
	prefixLength := flag.Int("l", 2, "prefix length (default: 2)")
	showHelpFlag := flag.Bool("help", false, "show help")
	flag.Parse()

	if *showHelpFlag {
		showHelp()
		return
	}

	// Validate max words (-w)
	if *maxWords <= 0 {
		err = errors.New("max words (-w) must be a positive integer")
		return
	}
	if *maxWords > 10000 {
		err = errors.New("max words (-w) cannot exceed 10,000")
		return
	}

	// Validate prefix length (-l)
	if *prefixLength <= 0 {
		err = errors.New("prefix length (-l) must be greater than 0")
		return
	}

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Read input
	var text string
	text, err = readInput()
	if err != nil {
		return
	}

	// Initialize Markov Chain
	markovChain := NewMarkovChain()

	// Build the chain
	err = markovChain.Build(text, *prefixLength)
	if err != nil {
		return
	}

	// Generate and print random text
	var generatedText string
	generatedText, err = markovChain.Generate(text, *startingPrefix, *prefixLength, *maxWords)
	if err != nil {
		return
	}

	fmt.Println(generatedText)
}
