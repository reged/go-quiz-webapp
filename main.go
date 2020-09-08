package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/hackersandslackers/golang-helloworld/internal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var wp *internal.WikiPage

func registerPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/register.html")
	if err != nil {
		log.Println("Error parse template :", err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/login.html")
	if err != nil {
		log.Println("Error parse template :", err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	userpassword := r.FormValue("userpassword")
	_, err := internal.CreateUser(db, username, userpassword)
	if err != nil {
		log.Printf("Registering new user failed: %s", err.Error())
	} else {
		log.Printf("User \"%s\" registered succefully", username)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	userpassword := r.FormValue("userpassword")
	result, err := internal.CheckUserPassword(db, username, userpassword)
	if err != nil {
		panic(err.Error())
	}
	if result {
		log.Printf("User \"%s\" logged in succesfully", username)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		log.Printf("User \"%s\" log in attemp failed", username)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func rebuildHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
	log.Print("Rebuild started...")
	_ = exec.Command("/bin/bash", "-x ./rebuild.sh")
	log.Print("Rebuilded")
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", loginPageHandler).Methods("GET")
	r.HandleFunc("/register", registerPageHandler).Methods("GET")
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginPageHandler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/rebuild", rebuildHandler).Methods("GET")
	fileServer := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(http.StripPrefix("/static", fileServer))
	return r
}

// Initiate web server
func main() {

	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Server started\n")
	db = internal.CreateConn("quiz_admin", "123", "quiz")
	// wp = GetWikiPage(db, "test")
	// _ = UpdateWikiPage(db, "test", []byte("Hello"))

	log.Fatal(srv.ListenAndServe())
}
