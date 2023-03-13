package main

import (
	"fmt"
	"math/rand"
	"time"
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
}

func main() {

	rand.Seed(time.Now().UnixNano())

	tagetWord := randomWord()
	fmt.Println(tagetWord)

}

func randomWord() string {
	tagetWord := dictionary[rand.Intn(len(dictionary))]
	return tagetWord
}
