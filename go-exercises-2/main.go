package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var (
	router *mux.Router
	db     map[string]user
	//Key is UUID, value is email
	sessions map[string]string
)

type user struct {
	First    string
	password []byte
}

func handleRequests() {
	router = mux.NewRouter().StrictSlash(true)
	db = make(map[string]user)
	sessions = make(map[string]string)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/register", register).Methods("POST")
	router.HandleFunc("/login", login)
	router.HandleFunc("/logout", logout).Methods("POST")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	indexPage, err := os.ReadFile("index.html")

	if err != nil {
		fmt.Println("Index page could not be loaded")
	}

	msg := r.FormValue("msg")

	sidCookieName := "sessionID"

	c, err := r.Cookie(sidCookieName)
	if err != nil {
		c = &http.Cookie{
			Name:  sidCookieName,
			Value: "",
		}
	}

	sid, err := parseToken(c.Value)
	if err != nil {
		fmt.Println(err.Error())
	}

	var email string
	if sid != "" {
		email = sessions[sid]
	}

	var first string
	if user, ok := db[email]; ok {
		first = user.First
	}

	fmt.Fprintf(w, string(indexPage), first, email, msg)
}

func register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("e")
	if email == "" {
		msg := url.QueryEscape("your email needs to not be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	if _, ok := db[email]; ok {
		msg := url.QueryEscape("email is used")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	password := r.FormValue("p")
	if password == "" {
		msg := url.QueryEscape("your password needs to not be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	bsp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		msg := url.QueryEscape("Bcrypt internal server error")
		http.Redirect(w, r, "/?msg="+msg, http.StatusInternalServerError)
		return
	}

	firstName := r.FormValue("first")
	if firstName == "" {
		msg := url.QueryEscape("your first name needs to not be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	db[email] = user{
		password: bsp,
		First:    firstName,
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("e")
	if email == "" {
		msg := url.QueryEscape("your email needs to not be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	password := r.FormValue("p")
	if password == "" {
		msg := url.QueryEscape("your password needs to not be empty")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	user, ok := db[email]
	if !ok {
		msg := url.QueryEscape("Your email or password doesn't match")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	err := bcrypt.CompareHashAndPassword(user.password, []byte(password))
	if err != nil {
		msg := url.QueryEscape("Your email or password doesn't match")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	addSession(email, w, r)

	msg := url.QueryEscape("you logged in: " + email)
	http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
}

func addSession(email string, w http.ResponseWriter, r *http.Request) {
	sid := createSessionId()
	sessions[sid] = email
	token, err := createToken(sid)
	if err != nil {
		msg := url.QueryEscape("couldn't create a token")
		http.Redirect(w, r, "/?msg="+msg, http.StatusSeeOther)
		return
	}

	c := http.Cookie{
		Name:  "sessionID",
		Value: token,
	}

	http.SetCookie(w, &c)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("sessionID")
	if err != nil {
		fmt.Println("cookie not found on logout")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sid, err := parseToken(c.Value)
	if err != nil {
		log.Panicln("logout parseToken")
	}

	delete(sessions, sid)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createSessionId() string {
	uuid := uuid.New().String()
	return uuid
}

func main() {
	handleRequests()
}
