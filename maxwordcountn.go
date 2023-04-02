package piscine

import (
	"sort"
)

// MaxWordCountN returns a map of the n most frequently occurring words in the input string 'text',
// along with their respective counts.
func MaxWordCountN(text string, n int) map[string]int {
	// Split 'text' into individual words using spaces as delimiters
	words := make([]string, 0)
	start := 0
	for i, r := range text {
		if r == ' ' {
			words = append(words, text[start:i])
			start = i + 1
		}
	}
	words = append(words, text[start:])

	// Count the occurrences of each word in 'words'
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// Sort the words in descending order of frequency and ascending order of word (for ties)
	sortedWords := make([]string, 0, len(wordCount))
	for word := range wordCount {
		sortedWords = append(sortedWords, word)
	}
	sort.Slice(sortedWords, func(i, j int) bool {
		if wordCount[sortedWords[i]] != wordCount[sortedWords[j]] {
			return wordCount[sortedWords[i]] > wordCount[sortedWords[j]]
		}
		return sortedWords[i] < sortedWords[j]
	})

	// Restrict the results to the top 'n' words (or less if there are ties)
	if n > len(sortedWords) {
		n = len(sortedWords)
	}
	tiedWords := sortedWords[:n]

	// Construct a new map containing the top 'n' words and their counts
	result := make(map[string]int)
	for _, word := range tiedWords {
		result[word] = wordCount[word]
	}
	return result
}
