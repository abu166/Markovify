package markov

//
//import (
//	"fmt"
//	"log"
//)
//
//// HandleError handles any errors that occur during the execution
//func HandleError(err error, message string) {
//	if err != nil {
//		log.Fatalf("Error: %s: %v", message, err)
//	}
//}
//
//// ValidateFlags validates the flag values
//func ValidateFlags(maxWords, prefixLength int) error {
//	if maxWords < 1 || maxWords > 10000 {
//		return fmt.Errorf("number of words must be between 1 and 10000")
//	}
//	if prefixLength < 1 || prefixLength > 5 {
//		return fmt.Errorf("prefix length must be between 1 and 5")
//	}
//	return nil
//}
//
//// ValidatePrefix validates the starting prefix
//func ValidatePrefix(prefix string, model map[string][]string) error {
//	if prefix != "" && len(model) > 0 {
//		if _, exists := model[prefix]; !exists {
//			return fmt.Errorf("prefix '%s' not found in the input text", prefix)
//		}
//	}
//	return nil
//}
//
//// CheckInputText checks if there is input text
//func CheckInputText(words []string) error {
//	if len(words) == 0 {
//		return fmt.Errorf("no input text")
//	}
//	return nil
//}
