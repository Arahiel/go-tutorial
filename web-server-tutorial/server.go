package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var router *mux.Router
var counter int
var mutex sync.Mutex

func init() {
	router = mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	router.HandleFunc("/increment", incrementHandler)
}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	log.Fatal(http.ListenAndServe(":8081", router))
}
