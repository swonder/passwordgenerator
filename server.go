package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	response := Response{
		NumberPasswords: 1,
		PasswordSize:    14,
		SpecChars:       "checked",
		UcChars:         "checked",
		Passwords:       []string{},
	}

	m := martini.Classic()

	m.Use(martini.Static("assets"))
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "home", response)
		response.Passwords = []string{}
	})

	m.Post("/", func(req *http.Request, r render.Render) {
		gatherInput(req, &response)

		validateInput(&response)

		r.Redirect("/")
	})

	m.Run()
}

func gatherInput(req *http.Request, res *Response) {
	numberPasswords, err := strconv.Atoi(req.FormValue("number"))
	if err != nil {
		log.Print("Unable to convert form number of passwords variable to a number")
	}
	res.NumberPasswords = numberPasswords

	passwordSize, err := strconv.Atoi(req.FormValue("size"))
	if err != nil {
		log.Print("Unable to convert form password size input to a number")
	}
	res.PasswordSize = passwordSize

	formSpecChars := req.FormValue("specchars")
	res.SpecChars = ""
	if formSpecChars == "on" {
		res.SpecChars = "checked"
	}

	formUcChars := req.FormValue("ucchars")
	res.UcChars = ""
	if formUcChars == "on" {
		res.UcChars = "checked"
	}
}

func validateInput(res *Response) {
	validationError := false

	if res.NumberPasswords < 1 || res.NumberPasswords > 100 {
		validationError = true
		res.Passwords = append(res.Passwords, "Number of passwords must be a number more than 0 and no more than 100")
	}

	if res.PasswordSize < 1 || res.PasswordSize > 75 {
		validationError = true
		res.Passwords = append(res.Passwords, "Password lengths must be a number more than 0 and no more than 75")
	}

	if !validationError {
		for i := 0; i < res.NumberPasswords; i++ {
			res.Passwords = append(res.Passwords, generateRandomString(res.PasswordSize, res.SpecChars, res.UcChars))
		}
	}
}
