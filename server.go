package main

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "gopkg.in/unrolled/render.v1"
    "net/http"
)

var printer = render.New(render.Options {
    Layout: "layout",
    IndentJSON: true,
})

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.Path("/").HandlerFunc(HomeHandler).Name("home")

    urls := router.PathPrefix("/urls").Subrouter()
    urls.Methods("GET").Path("/").HandlerFunc(UrlsIndexHandler).Name("urls_index")

    app := negroni.Classic()
    app.UseHandler(router)
    app.Run(":8080")
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
    data := map[string]interface{} {
        "title": "Home",
    }
    printer.HTML(w, http.StatusOK, "home", data)
}

func UrlsIndexHandler(w http.ResponseWriter, req *http.Request) {
    urls := []string {
        "www.google.com",
        "www.yahoo.com",
        "www.cnn.com",
    }
    printer.JSON(w, http.StatusOK, map[string][]string {"urls": urls})
}
