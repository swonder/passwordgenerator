package main

import (
  //"log"
  "strconv"
  "net/http"
  "math/rand"
)

import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"


type Response struct {
  Number string
  Size string
  Specchars string
  Ucchars string
  Passwords []string
}

var ucletters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lcletters = []rune("abcdefghijklmnopqrstuvwxyz")
var ucchars =   []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*.-|")
var lcchars =   []rune("abcdefghijklmnopqrstuvwxyz!@#$%^&*.-|")

//Create a random length string
func randString(n int, special string, uc string) string {
    b := make([]rune, n)
    if special == "checked" && uc == "checked" {
      for i := range b {
          b[i] = ucchars[rand.Intn(len(ucchars))]
      }
    } else if special != "checked" && uc == "checked" {
      for i := range b {
          b[i] = ucletters[rand.Intn(len(ucletters))]
      }
    } else if special == "checked" && uc != "checked" {
      for i := range b {
          b[i] = lcchars[rand.Intn(len(lcchars))]
      }
    } else {
      for i := range b {
          b[i] = lcletters[rand.Intn(len(lcletters))]
      }
    }

    return string(b)
}

func main() {
  var pwds []string

  response := Response {
    Number:             "1",
    Size:               "14",
    Specchars:          "checked",
    Ucchars:            "checked",
    Passwords:          []string{},
  }

  m := martini.Classic()
  //Include css, js, etc
  m.Use(martini.Static("assets"))
  m.Use(render.Renderer())

  // render.Options{Layout: "layout",

  //Root route
  m.Get("/", func(r render.Render) {
       r.HTML(200, "home", response)
       response.Passwords = []string{}
  })

  //Posting to a form in root route
  m.Post("/" ,func(req *http.Request, r render.Render) {
      pwds = []string{}
      formnumber := req.FormValue("number")
      formsize := req.FormValue("size")
      formspecchars := req.FormValue("specchars")
      formucchars := req.FormValue("ucchars")
      if formspecchars == "on" {
        formspecchars = "checked"
      }
      if formucchars == "on" {
        formucchars = "checked"
      }

      if len(formsize) > 0 {
        number, err := strconv.Atoi(formnumber)
        size, err := strconv.Atoi(formsize)
        if number > 0 && number < 100 {
          if err == nil {
            if size > 0 && size <= 75 {
              //Loop through and create 'number' number of passwords
              for i := 0; i < number; i++ {
                pwds = append(pwds, randString(size, formspecchars, formucchars))
                //strconv.Itoa(i+1) + ":  " +
              }

              response = Response {
                Number:             strconv.Itoa(number),
                Size:               strconv.Itoa(size),
                Specchars:          formspecchars,
                Ucchars:            formucchars,
                Passwords:          pwds,
              }

            } else {
              pwds = append([]string{}, "Password length must be greater than 0 and less than 75 characters ")
              response.Passwords = pwds
            }
          } else {
            pwds = append([]string{}, "Password length must be a number")
            response.Passwords = pwds
          }
        } else {
          pwds = append([]string{}, "Number of passwords must be greater than 0 and less or equal to 100")
          response.Passwords = pwds
        }

      } else {
        pwds = append([]string{}, "Please make sure number of passwords and password lengths are set")
        response.Passwords = pwds
      }
      r.Redirect("/")
  })

  m.Run()
}
