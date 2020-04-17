package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

var myStoredArticles Articles //A stored collection of all our articles

func allArticles(w http.ResponseWriter, r *http.Request) {
	myStoredArticles = append(myStoredArticles, Article{Id: "1", Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"})

	testArticle := Article{Id: "3", Title: "Fun Times", Desc: "Test Fun!", Content: "Shutup world"}
	myStoredArticles = append(myStoredArticles, testArticle)

	fmt.Println("Endpoint Hit: All Articles Endpoint. This is what you get with GET")
	json.NewEncoder(w).Encode(myStoredArticles)
}
func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Printf("The key is: %v\n", key)

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range myStoredArticles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Here's our body: \n%v\n", reqBody)
	var newArticle Article
	json.Unmarshal(reqBody, &newArticle)
	myStoredArticles = append(myStoredArticles, newArticle)
	json.NewEncoder(w).Encode(newArticle)
	fmt.Printf("Here's our Json: \n%v\n", newArticle)
	//fmt.Fprintf(w, "%+v", string(reqBody))
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range myStoredArticles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			myStoredArticles = append(myStoredArticles[:index], myStoredArticles[index+1:]...)
		}
	}

}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Here's our body: \n%v\n", reqBody)
	var updatedArticle Article
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range myStoredArticles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to update the
			// article
			/*
				newArticle := Article{
					Id:      article.Id,
					Title:   "Our updated Title field!",
					Desc:    "Our updated Description!",
					Content: "Our updated Content!",
				}
			*/
			myStoredArticles[index] = updatedArticle
		}
	}
}

//Test func
func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello from sayHi!\n")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Here's our body: \n%v\n", reqBody)
	var newArticle Article
	json.Unmarshal(reqBody, &newArticle)
	fmt.Printf("Here is our newArticle: \n%v\n", newArticle)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	myRouter.HandleFunc("/sayhi", sayHi).Methods("POST")
	myRouter.HandleFunc("/postArticle", updateArticle).Methods("POST")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST") //This has to be defined before the below!
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
