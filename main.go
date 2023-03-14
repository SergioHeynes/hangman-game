package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

var dictionary = []string{
	"Corona",
	"Caballero",
	"Quisco",
	"Septiembre",
	"Interesante",
	"Guitarra",
	"Mesopotamia",
	"Mural",
	"Abecedario",
	"Sistema solar",
	"Estados Unidos de América",
}

func main() {

	rand.Seed(time.Now().UnixNano())

	targetWoerd := randomWord()

	gessedLetters := initializeGuessedWords(targetWoerd)
	printGameState(targetWoerd, gessedLetters)

	// fmt.Println(tagetWord)

}

func initializeGuessedWords(targetWoerd string) map[rune]bool {
	gessedLetters := map[rune]bool{}
	gessedLetters[unicode.ToLower(rune(targetWoerd[0]))] = true
	gessedLetters[unicode.ToLower(rune(targetWoerd[len(targetWoerd)-1]))] = true

	return gessedLetters
}

func randomWord() string {
	tagetWord := dictionary[rand.Intn(len(dictionary))]
	return tagetWord
}

func printGameState(targetWord string, gessedLetters map[rune]bool) {

	for _, ch := range targetWord {

		if ch == ' ' {
			fmt.Print(" ")
		} else if gessedLetters[unicode.ToLower(ch)] {

			fmt.Print(string(ch))

		} else {
			fmt.Print("_")
		}

		fmt.Print(" ")

	}
	fmt.Println()

}
