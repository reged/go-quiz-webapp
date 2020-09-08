package controller

import (
	"fmt"
	"github.com/hackersandslackers/golang-helloworld/internal/database"
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

func (c Controller) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
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

func (c Controller) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
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

func (c Controller) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	userName := r.FormValue("username")
	userPassword := r.FormValue("userpassword")
	_, err := database.CreateUser(c.Db, userName, userPassword)
	if err != nil {
		log.Printf("Registering new user failed: %s", err.Error())
	} else {
		log.Printf("User \"%s\" registered succefully", userName)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (c Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	userPassword := r.FormValue("userpassword")
	result, err := database.CheckUserPassword(c.Db, userName, userPassword)
	if err != nil {
		panic(err.Error())
	}
	if result {
		log.Printf("User \"%s\" logged in succesfully", userName)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		log.Printf("User \"%s\" log in attemp failed", userName)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (c Controller) RebuildHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Rebuild started...")
	_ = exec.Command("/bin/bash", "-x ./rebuild.sh")
	log.Print("Rebuilded")
}

func (c Controller) TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello\n")
}
