package markov

//import (
//	"fmt"
//	"os"
//	// "math/rand"
//	// "time"
//)
//
//// PrintUsage prints the usage of the program
//func PrintUsage() {
//	fmt.Println("Markov Chain text generator.")
//	fmt.Println("Usage:")
//	fmt.Println("  markovchain [-w <N>] [-p <S>] [-l <N>]")
//	fmt.Println("Options:")
//	fmt.Println("  --help  Show this screen.")
//	fmt.Println("  -w N    Number of maximum words")
//	fmt.Println("  -p S    Starting prefix")
//	fmt.Println("  -l N    Prefix length")
//}
//
//func HasInput() bool {
//	info, err := os.Stdin.Stat()
//	if err != nil {
//		// Error checking stdin state
//		return false
//	}
//
//	// Check if the input stream is a terminal (interactive) or if data is available
//	// If it is not a terminal and has data, we assume input is available.
//	return (info.Mode() & os.ModeCharDevice) == 0
//}

// Initialize random seed
// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }
