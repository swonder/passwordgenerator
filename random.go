package main

import (
	"crypto/rand"
	"errors"
	"log"
	"math/big"
)

var lcLetters = []rune("abcdefghijklmnopqrstuvwxyz")
var ucLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lcChars = []rune("abcdefghijklmnopqrstuvwxyz!@#$%^&*.-|")
var ucChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*.-|")

func getRandomChar(charset []rune) (char rune, e error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err == nil {
		return charset[index.Uint64()], nil
	}

	errorMsg := "an error has occurred generating a random character"
	log.Print(errorMsg)
	return ' ', errors.New(errorMsg)
}

func generateRandomString(n int, special string, uc string) string {
	b := make([]rune, n)
	var charSet []rune

	if special == "checked" && uc == "checked" {
		charSet = ucChars
	} else if special != "checked" && uc == "checked" {
		charSet = ucLetters
	} else if special == "checked" && uc != "checked" {
		charSet = lcChars
	} else {
		charSet = lcLetters
	}

	for i := range b {
		randChar, err := getRandomChar(charSet)
		if err == nil {
			b[i] = randChar
		} else {
			log.Fatal(err)
		}
	}

	return string(b)
}
