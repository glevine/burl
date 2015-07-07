package main

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.Path("/").HandlerFunc(HomeHandler).Name("home")

    resources := router.PathPrefix("/urls").Subrouter()
    resources.Methods("GET").Path("/").HandlerFunc(ResourcesIndexHandler).Name("resources_index")

    n := negroni.Classic()
    n.UseHandler(router)
    n.Run(":8080")
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
}

func ResourcesIndexHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Get me some urls!")
}
