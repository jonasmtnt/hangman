package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	playHangman()
}

func getWord() string {
	words := []string{"python", "javascript", "programming", "computer", "science"}
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func playHangman() {
	word := getWord()
	wordLetters := make(map[rune]bool)
	for _, char := range word {
		wordLetters[char] = true
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	usedLetters := make(map[rune]bool)
	lives := 6

	for len(wordLetters) > 0 && lives > 0 {
		fmt.Print("You have used these letters: ")
		for char := range usedLetters {
			fmt.Printf("%c ", char)
		}
		fmt.Println()

		var wordList []string
		for _, char := range word {
			if usedLetters[char] {
				wordList = append(wordList, string(char))
			} else {
				wordList = append(wordList, "-")
			}
		}
		fmt.Printf("Current word: %s\n", strings.Join(wordList, " "))

		var userLetter rune
		fmt.Print("Guess a letter: ")
		fmt.Scanf("%c\n", &userLetter)

		if strings.ContainsRune(alphabet, userLetter) && !usedLetters[userLetter] {
			usedLetters[userLetter] = true
			if wordLetters[userLetter] {
				delete(wordLetters, userLetter)
			} else {
				lives--
				fmt.Println("Letter is not in word.")
			}

		} else if usedLetters[userLetter] {
			fmt.Println("You have already used that character. Please try again.")
		} else {
			fmt.Println("Invalid character. Please try again.")
		}
	}

	if lives == 0 {
		fmt.Printf("You died, sorry. The word was %s\n", word)
	} else {
		fmt.Printf("You win! The word was %s\n", word)
	}
}
