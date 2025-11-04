package markov

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//	"math/rand"
//)
//
//
//// ReadInputAndBuildModel reads the input text and builds the Markov chain model
//func ReadInputAndBuildModel(markov *MarkovChain) error {
//	// Open stdin and read the entire input
//	scanner := bufio.NewScanner(os.Stdin)
//	var words []string
//
//	// Read input text and split it into words
//	for scanner.Scan() {
//		line := scanner.Text()
//		// If the line is not empty, add its words to the word list
//		words = append(words, strings.Fields(line)...)
//	}
//
//	// Handle any scanner errors
//	if err := scanner.Err(); err != nil {
//		return fmt.Errorf("Error reading input: %v", err)
//	}
//
//	// Check if input is empty (no text was read)
//	if len(words) == 0 {
//		return fmt.Errorf("no input text")
//	}
//
//	// Build the Markov model using word sequences
//	for i := 0; i <= len(words)-markov.PrefixLength; i++ {
//		prefix := strings.Join(words[i:i+markov.PrefixLength], " ")
//		if i+markov.PrefixLength < len(words) {
//			suffix := words[i+markov.PrefixLength]
//			markov.Model[prefix] = append(markov.Model[prefix], suffix)
//		}
//	}
//
//	return nil
//}
//
//// ValidatePrefix checks if the provided prefix exists in the model
//func (markov *MarkovChain) ValidatePrefix(prefix string) bool {
//	_, exists := markov.Model[prefix]
//	return exists
//}
//
//
//// GenerateText generates random text based on the Markov model
//func GenerateText(markov *MarkovChain, prefix string, maxWords int) error {
//	// Validate prefix if provided
//	if err := ValidatePrefix(prefix, markov.Model); err != nil {
//		return err
//	}
//
//	// Start with the provided prefix, or the first available prefix
//	if prefix == "" {
//		for k := range markov.Model {
//			prefix = k
//			break
//		}
//	}
//
//	// Generate words
//	wordCount := 0
//	currentPrefix := prefix
//
//	// Print the initial prefix
//	fmt.Print(currentPrefix + " ")
//
//	for wordCount < maxWords {
//		suffixes, exists := markov.Model[currentPrefix]
//		if !exists {
//			break
//		}
//
//		// Pick a random suffix
//		randomIndex := rand.Intn(len(suffixes))
//		nextWord := suffixes[randomIndex]
//		fmt.Print(nextWord + " ")
//
//		// Update the prefix to be the last two words
//		wordsInPrefix := strings.Fields(currentPrefix)
//		if len(wordsInPrefix) > 1 {
//			currentPrefix = wordsInPrefix[1] + " " + nextWord
//		} else {
//			currentPrefix = nextWord
//		}
//
//		wordCount++
//	}
//	fmt.Println()
//
//	return nil
//}
