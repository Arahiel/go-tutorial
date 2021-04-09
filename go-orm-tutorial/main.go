package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connectionString = "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
var driverName = "postgres"

type User struct {
	gorm.Model
	Name  string
	Email string
}

func initialMigration() {
	db, err := gorm.Open(driverName, connectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", allUsersGetHandler).Methods("GET")
	router.HandleFunc("/user/{name}", userDeleteHandler).Methods("DELETE")
	router.HandleFunc("/user/{name}/{email}", userPostHandler).Methods("POST")
	router.HandleFunc("/user/{name}/{email}", userPutHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func allUsersGetHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(driverName, connectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println(users)

	json.NewEncoder(w).Encode(users)
}

func userPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	email := vars["email"]

	db, err := gorm.Open(driverName, connectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New user created successfully")
}

func userPutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	email := vars["email"]

	db, err := gorm.Open(driverName, connectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.Save(&user)

	fmt.Fprintf(w, "User updated successfully")
}

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	db, err := gorm.Open(driverName, connectionString)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User deleted successfully")
}

func main() {
	fmt.Println("Go ORM Tutorial")

	initialMigration()
	handleRequests()
}
