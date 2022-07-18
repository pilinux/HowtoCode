// a set contains unique values

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a search word, exiting...")
		return
	}
	query := args[0]
	queryToLower := strings.ToLower(query)

	// find only letters, omit everything else
	rx := regexp.MustCompile(`[^a-z]+`)

	in := bufio.NewScanner(os.Stdin)
	wordCount := 0
	uniqueWordCount := 0

	// unique word will be key, value is boolean
	// keys are case-sensitive
	words := make(map[string]bool)

	// scan word by word
	in.Split(bufio.ScanWords)

	for in.Scan() {
		wordCount++
		// fmt.Println(in.Text())

		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")
		// fmt.Println(word)

		// index a word if it has min 3 letters
		if len(word) > 2 {
			if !words[word] {
				words[word] = true
				uniqueWordCount++
			}
		}

	}

	fmt.Println("Total words: ", wordCount)
	fmt.Println("Total unique words: ", uniqueWordCount)

	if words[queryToLower] {
		fmt.Println("The input contains the complete word: ", query)
		return
	}

	fmt.Println("The input does not contain the complete word: ", query)
}

// ./mapASset spacex < mapASset.txt
// curl -s https://raw.githubusercontent.com/piLinuxME/piLinux.me/master/content/golang/slice.md | ./mapASset program
