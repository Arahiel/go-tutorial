package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
	mutex  sync.Mutex
)

type Article struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

var Articles []Article

func init() {
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/favicon.ico", faviconGetHandler)
	router.HandleFunc("/articles", articlesGetHandler).Methods("GET")
	router.HandleFunc("/articles", articleCreateHandler).Methods("POST")
	router.HandleFunc("/articles", articleUpdateHandler).Methods("PUT")
	router.HandleFunc("/articles/{id}", articleDeleteHandler).Methods("DELETE")
	router.HandleFunc("/articles/{id}", articleGetHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home hit!")
	w.WriteHeader(200)
}

func articlesGetHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "	")
	encoder.Encode(Articles)
}

func articleGetHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, article := range Articles {
		if article.Id == id {
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "	")
			encoder.Encode(article)
			return
		}
	}
	w.WriteHeader(404)
}

func articleCreateHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	byteContent, _ := io.ReadAll(r.Body)
	var article Article
	json.Unmarshal(byteContent, &article)
	Articles = append(Articles, article)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "	")
	encoder.Encode(article)
	mutex.Unlock()
}

func articleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	id := mux.Vars(r)["id"]
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			mutex.Unlock()
			return
		}
	}
	mutex.Unlock()
	w.WriteHeader(404)
}

func articleUpdateHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	byteContent, _ := io.ReadAll(r.Body)
	var newArticle Article
	err := json.Unmarshal(byteContent, &newArticle)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "	")

	if err != nil {
		encoder.Encode(err)
		w.WriteHeader(400)
		mutex.Unlock()
		return
	}

	for index, article := range Articles {
		if article.Id == newArticle.Id {
			Articles[index] = newArticle
			encoder.Encode(newArticle)
			mutex.Unlock()
			return
		}
	}
	mutex.Unlock()
	w.WriteHeader(404)
}

func faviconGetHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "resources/favicon.ico")
}

func main() {
	Articles = []Article{
		{"1", "Title 1", "Desc 1", "Content 1"},
		{"2", "Title 2", "Desc 2", "Content 2"},
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}
