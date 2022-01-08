package main

import (
	"math/rand"
	"testing"
	"time"
)

func setup() {
	rand.Seed(time.Now().UnixNano())
}

func TestGetRandomChar(t *testing.T) {
	charSets := [][]rune{lcLetters, ucLetters, lcChars, ucChars}

	for _, charSet := range charSets {
		char, err := getRandomChar(charSet)
		charFound := false

		for _, val := range charSet {
			if val == char {
				charFound = true
			}
		}

		if !charFound {
			t.Log("Rune returned from getRandomChar() was not found in original slice")
			t.Fail()
		}

		if err != nil {
			t.Log("Error was returned from function when a rune should have been returned")
			t.Fail()
		}
	}
}

func TestGenerateRandomStringWithNoUpperCaseOrSpecChars(t *testing.T) {
	setup()

	stringSize := rand.Intn(75-1) + 1

	result := generateRandomString(stringSize, "", "")

	if len(result) != stringSize {
		t.Log("Size of string generated does not match input size")
		t.Fail()
	}

	charNotFoundInResult := false
	for _, returnedChar := range result {
		charFound := false
		for _, validChar := range lcLetters {
			if returnedChar == validChar {
				charFound = true
			}
		}
		if !charFound {
			charNotFoundInResult = true
		}
	}

	if charNotFoundInResult {
		t.Log("A character in returned string was not found in list of valid characters (Lowercase characters")
		t.Fail()
	}
}

func TestGenerateRandomStringWithUpperAndLowerCaseChars(t *testing.T) {
	setup()

	stringSize := rand.Intn(75-1) + 1

	result := generateRandomString(stringSize, "", "checked")

	if len(result) != stringSize {
		t.Log("Size of string generated does not match input size")
		t.Fail()
	}

	charNotFoundInResult := false
	for _, returnedChar := range result {
		charFound := false
		for _, validChar := range ucLetters {
			if returnedChar == validChar {
				charFound = true
			}
		}
		if !charFound {
			charNotFoundInResult = true
		}
	}

	if charNotFoundInResult {
		errMsg := `A character in returned string was not found in list of valid characters
			(Uppercase and lowercase characters)`
		t.Log(errMsg)
		t.Fail()
	}
}

func TestGenerateRandomStringWithLowerCaseAndSpecChars(t *testing.T) {
	setup()

	stringSize := rand.Intn(75-1) + 1

	result := generateRandomString(stringSize, "checked", "")

	if len(result) != stringSize {
		t.Log("Size of string generated does not match input size")
		t.Fail()
	}

	charNotFoundInResult := false
	for _, returnedChar := range result {
		charFound := false
		for _, validChar := range lcChars {
			if returnedChar == validChar {
				charFound = true
			}
		}
		if !charFound {
			charNotFoundInResult = true
		}
	}

	if charNotFoundInResult {
		errMsg := `A character in returned string was not found in list of valid characters
			(Lowercase and special characters)`
		t.Log(errMsg)
		t.Fail()
	}
}

func TestGenerateRandomStringWithAllChars(t *testing.T) {
	setup()

	stringSize := rand.Intn(75-1) + 1

	result := generateRandomString(stringSize, "checked", "checked")

	if len(result) != stringSize {
		t.Log("Size of string generated does not match input size")
		t.Fail()
	}

	charNotFoundInResult := false
	for _, returnedChar := range result {
		charFound := false
		for _, validChar := range ucChars {
			if returnedChar == validChar {
				charFound = true
			}
		}
		if !charFound {
			charNotFoundInResult = true
		}
	}

	if charNotFoundInResult {
		errMsg := `A character in returned string was not found in list of valid characters
			(Uppercase, lowercase, and special characters)`
		t.Log(errMsg)
		t.Fail()
	}
}
