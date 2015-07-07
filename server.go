package main

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
      fmt.Fprintf(w, "Welcome to the home page!")
    })

    n := negroni.Classic()
    n.UseHandler(router)
    n.Run(":8080")
}
