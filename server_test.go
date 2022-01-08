package main

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGatherInputSetsPasswordNumberResponseVarCorrectly(t *testing.T) {
	res := Response{}
	formData := url.Values{
		"number":    {"10"},
		"size":      {"0"},
		"specchars": {""},
		"ucchars":   {""},
	}
	req, err := http.NewRequest("POST", "http://_", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Log("There was an error creating the request")
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	gatherInput(req, &res)

	if res.NumberPasswords != 10 {
		t.Log("Expected number of passwords is incorrect in response struct")
		t.Fail()
	}
}

func TestGatherInputSetsPasswordResponseVarCorrectly(t *testing.T) {
	res := Response{}
	formData := url.Values{
		"number":    {"0"},
		"size":      {"20"},
		"specchars": {""},
		"ucchars":   {""},
	}
	req, err := http.NewRequest("POST", "http://_", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Log("There was an error creating the request")
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	gatherInput(req, &res)

	if res.PasswordSize != 20 {
		t.Log("Expected size of each password is incorrect in response struct")
		t.Fail()
	}
}

func TestGatherInputSetsSpecialCharsResponseVarCorrectly(t *testing.T) {
	res := Response{}
	formData := url.Values{
		"number":    {"0"},
		"size":      {"0"},
		"specchars": {"on"},
		"ucchars":   {""},
	}
	req, err := http.NewRequest("POST", "http://_", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Log("There was an error creating the request")
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	gatherInput(req, &res)

	if res.SpecChars != "checked" {
		t.Log("SpecChars value should be 'checked' in response struct")
		t.Fail()
	}

	// Test unchecked value
	res = Response{}
	formData = url.Values{
		"number":    {"0"},
		"size":      {"0"},
		"specchars": {""},
		"ucchars":   {""},
	}
	req, err = http.NewRequest("POST", "http://_", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Log("There was an error creating the request")
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	gatherInput(req, &res)

	if res.SpecChars != "" {
		t.Log("SpecChars value should be empty in response struct")
		t.Fail()
	}
}

func TestGatherInputSetsUcCharsResponseVarCorrectly(t *testing.T) {
	res := Response{}
	formData := url.Values{
		"number":    {"0"},
		"size":      {"0"},
		"specchars": {""},
		"ucchars":   {"on"},
	}
	req, err := http.NewRequest("POST", "http://_", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Log("There was an error creating the request")
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	gatherInput(req, &res)

	if res.UcChars != "checked" {
		t.Log("UcChars value should be 'checked' in response struct")
		t.Fail()
	}

	// Test unchecked value
	res = Response{}
	formData = url.Values{
		"number":    {"0"},
		"size":      {"0"},
		"specchars": {""},
		"ucchars":   {""},
	}
	req, err = http.NewRequest("POST", "http://_", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Log("There was an error creating the request")
		t.Fail()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	gatherInput(req, &res)

	if res.UcChars != "" {
		t.Log("UcChars value should be empty in response struct")
		t.Fail()
	}
}

func TestValidateInputFailsTooManyPasswords(t *testing.T) {
	response := Response{
		NumberPasswords: 500,
		PasswordSize:    20,
		SpecChars:       "checked",
		UcChars:         "checked",
		Passwords:       []string{},
	}

	validateInput(&response)

	errorFound := false
	for _, v := range response.Passwords {
		if v == "Number of passwords must be a number more than 0 and no more than 100" {
			errorFound = true
		}
	}

	if !errorFound {
		t.Log("Expected error was not found")
		t.Fail()
	}
}

func TestValidateInputFailsNotEnoughPasswords(t *testing.T) {
	response := Response{
		NumberPasswords: 0,
		PasswordSize:    20,
		SpecChars:       "checked",
		UcChars:         "checked",
		Passwords:       []string{},
	}

	validateInput(&response)

	errorFound := false
	for _, v := range response.Passwords {
		if v == "Number of passwords must be a number more than 0 and no more than 100" {
			errorFound = true
		}
	}

	if !errorFound {
		t.Log("Expected error was not found")
		t.Fail()
	}
}

func TestValidateInputFailsPasswordSizeTooLong(t *testing.T) {
	response := Response{
		NumberPasswords: 1,
		PasswordSize:    76,
		SpecChars:       "checked",
		UcChars:         "checked",
		Passwords:       []string{},
	}

	validateInput(&response)

	errorFound := false
	for _, v := range response.Passwords {
		if v == "Password lengths must be a number more than 0 and no more than 75" {
			errorFound = true
		}
	}

	if !errorFound {
		t.Log("Expected error was not found")
		t.Fail()
	}
}

func TestValidateInputFailsPasswordSizeTooShort(t *testing.T) {
	response := Response{
		NumberPasswords: 1,
		PasswordSize:    0,
		SpecChars:       "checked",
		UcChars:         "checked",
		Passwords:       []string{},
	}

	validateInput(&response)

	errorFound := false
	for _, v := range response.Passwords {
		if v == "Password lengths must be a number more than 0 and no more than 75" {
			errorFound = true
		}
	}

	if !errorFound {
		t.Log("Expected error was not found")
		t.Fail()
	}
}
