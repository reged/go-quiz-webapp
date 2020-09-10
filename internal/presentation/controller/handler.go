package controller

import (
	"fmt"
	"github.com/reged/go-quiz-webapp/internal/database"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func (c Controller) IndexPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/index.html")
	cookie, err := r.Cookie("QuizUser")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Wrong Cookie :", err)
		return
	}
	err =  t.Execute(w, template.HTML(cookie.Value))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error parse template :", err)
		return
	}
}

func (c Controller) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/register.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error parse template :", err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error executing template :", err)
		return
	}
}

func (c Controller) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/login.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error parse template :", err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error executing template :", err)
		return
	}
}

func (c Controller) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	userName := r.FormValue("username")
	userPassword := r.FormValue("userpassword")
	_, err := database.CreateUser(c.Db, userName, userPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Registering new user failed: %s", err.Error())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("User \"%s\" registered succefully", userName)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (c Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	userPassword := r.FormValue("userpassword")
	expire := time.Now().AddDate(0, 0, 1)

	result, err := database.CheckUserPassword(c.Db, userName, userPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error while logging in")
	}
	if result {
		cookie := &http.Cookie{
			Name:       "QuizUser",
			Value:      userName,
			Path:       "/",
			Domain:     "localhost",
			Expires:    expire,
			RawExpires: expire.Format(time.UnixDate),
			MaxAge:     86400,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   nil,
		}
		http.SetCookie(w, cookie)
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
