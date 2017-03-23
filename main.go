package main

import (
  "log"
  "net/http"

  "github.com/codegangsta/negroni"
  "github.com/briankeane/goCourse/common"
  "github.com/briankeane/goCourse/routers"
)

func main() {
  common.StartUp()
  r := routers.InitRoutes()

  n := negroni.Classic()
  n.UseHandler(r)

  log.Println("Listening...")
  http.ListenAndServe(":8080", n)
}