package main

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "gopkg.in/unrolled/render.v1"
    "github.com/jmcvetta/neoism"
    "net/http"
    "os"
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

type Url struct {
    Node neoism.Node `json:"url"`
}

func UrlsIndexHandler(w http.ResponseWriter, req *http.Request) {
    neo, err := neoism.Connect(os.Getenv("GRAPHENEDB_URL"))

    if err != nil {
        printer.JSON(w, http.StatusServiceUnavailable, map[string]interface{} {
            "code": http.StatusServiceUnavailable,
            "error": http.StatusText(http.StatusServiceUnavailable),
            "message": err,
        })
    } else {
        urls := []Url{}
        cq := neoism.CypherQuery {
            Statement: `MATCH (url:Url) RETURN url`,
            Result: &urls,
        }
        err := neo.Cypher(&cq)
        if err != nil {
            printer.JSON(w, http.StatusServiceUnavailable, map[string]interface{} {
                "code": http.StatusServiceUnavailable,
                "error": http.StatusText(http.StatusServiceUnavailable),
                "message": err,
            })
        }
        printer.JSON(w, http.StatusOK, map[string][]Url {"urls": urls})
    }
}
