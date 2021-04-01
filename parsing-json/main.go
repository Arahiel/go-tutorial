package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var filename string

func init() {
	filename = "users.json"
}

type UsersDb struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error occurred during opening of the file")
		return
	}
	fmt.Printf("Successfully opened %s file", filename)
	defer jsonFile.Close()

	fileContent, _ := io.ReadAll(jsonFile)

	var db UsersDb
	json.Unmarshal(fileContent, &db)
	printUsers(db.Users)

}

func printUsers(users []User) {
	for _, user := range users {
		fmt.Println()
		fmt.Println("User Name: ", user.Name)
		fmt.Println("User Type: ", user.Type)
		fmt.Println("User Age: ", user.Age)
		fmt.Println("Social Facebook: ", user.Social.Facebook)
		fmt.Println("Social Twitter: ", user.Social.Twitter)
		fmt.Println()
	}
}
