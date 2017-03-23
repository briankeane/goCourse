package main

import (
  "github.com/taskmanager/common"
  "github.com/taskmanager/routers"

  "github.com/codegangsta/negroni"
  "github.com/briankeane/goCourse/common"
  "github.com/briankeane/goCourse/routers"
)

func main() {
  common.StartUp()
  r := routers.InitRoutes()

  n := negroni.Classic
  n.UseHandler(r)

  log.PrintLn("Listening...")
  http.ListenAndAserve(":8080", n)
}