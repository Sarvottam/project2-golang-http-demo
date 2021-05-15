package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
var port int = 10000

type Article struct {
	Id      string `json:"Id"`
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    // http.HandleFunc("/", homePage)
    // http.HandleFunc("/articles", returnAllArticles)
	 
	// creates a new instance of a mux router
	 myRouter := mux.NewRouter().StrictSlash(true)
	 // replace http.HandleFunc with myRouter.HandleFunc
	 myRouter.HandleFunc("/", homePage)
	 myRouter.HandleFunc("/all", returnAllArticles)
	 myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	 log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
    key := vars["id"]

    fmt.Fprintf(w, "Key: " + key)
	fmt.Println("Endpoint Hit: returnAllArticles",key)

    json.NewEncoder(w).Encode(Articles)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func main() {
	fmt.Println("Application running on port ",port)
	fmt.Println("Rest API v2.0 - Mux Routers")

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }

    handleRequests()
}