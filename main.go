package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	passphraseLen := getCommandLineArgs()

	codeToWord := loadWordlist()
	randomCodes := generateArrayOfRandomCodes(passphraseLen)

	var words []string

	for _, code := range randomCodes {
		words = append(words, codeToWord[code])
	}

	const wordsSeparator = " "

	passphrase := strings.Join(words, wordsSeparator)

	fmt.Printf("Your passphrase is: %s\n", passphrase)

	elapsed := time.Since(start)
	fmt.Printf("Function took %s\n", elapsed)
}

// Loads the wordlist from the file and returns a map of code to word
func loadWordlist() map[string]string {
	file, err := os.ReadFile("wordlist.txt")

	if err != nil {
		panic(err)
	}

	content := string(file)

	codeToWord := make(map[string]string)

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		code, word, _ := strings.Cut(line, "\t")
		codeToWord[code] = word
	}

	return codeToWord
}

// Parses the command line arguments and returns the number of words to generate
func getCommandLineArgs() (passphraseLen int) {
	const defaultPassphraseLen = 6

	args := os.Args[1:]

	if len(args) == 0 {
		return defaultPassphraseLen
	}

	numWords, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Please provide a valid number of words to generate")
		os.Exit(1)
	}

	return numWords
}

// Generates an array of random 6-digit codes
func generateArrayOfRandomCodes(len int) []string {
	var codes []string

	for i := 0; i < len; i++ {
		codes = append(codes, generateRandomCode(6))
	}

	return codes
}

// Generates a random x-digit code (1-6)
func generateRandomCode(len int) string {
	var code string

	for i := 0; i < len-1; i++ {
		randomInt := rand.Intn(5) + 1
		code += strconv.Itoa(randomInt)
	}

	return code
}
