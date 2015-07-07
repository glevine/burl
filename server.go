package main

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
    "net/http"
)

var printer = render.New(render.Options {
    IndentJSON: true,
})

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.Path("/").HandlerFunc(HomeHandler).Name("home")

    resources := router.PathPrefix("/urls").Subrouter()
    resources.Methods("GET").Path("/").HandlerFunc(ResourcesIndexHandler).Name("resources_index")

    app := negroni.Classic()
    app.UseHandler(router)
    app.Run(":8080")
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
    printer.JSON(w, http.StatusOK, map[string]string {"home": "Welcome to the home page!"})
}

func ResourcesIndexHandler(w http.ResponseWriter, req *http.Request) {
    urls := []string {
        "www.google.com",
        "www.yahoo.com",
        "www.cnn.com",
    }
    printer.JSON(w, http.StatusOK, map[string][]string {"urls": urls})
}
