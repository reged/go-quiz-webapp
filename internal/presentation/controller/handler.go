package controller

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/reged/go-quiz-webapp/internal/database"
)

func checkUserCookie(c Controller, username string, token string) (bool, error) {
	cookie, err := database.GetUserCookie(c.Db, username)
	if err != nil {
		return false, err
	}
	if cookie != token {
		return false, nil
	}
	return true, nil
}

func (c Controller) IndexPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/index.html")
	cookie, err := r.Cookie("QuizUser")
	if err != nil {
		log.Println("Wrong Cookie :", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	err = t.Execute(w, template.HTML(cookie.Value))
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

func generateToken() string {
	rand.Seed(time.Now().Unix())

	//Only lowercase
	charSet := "abcdedfghijklmnopqrst"
	var output strings.Builder
	length := 10
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
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
		token := generateToken()
		cookie = &http.Cookie{
			Name:       "QuizUserToken",
			Value:      token,
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
		err = database.SetUserCookie(c.Db, userName, token)
		if err != nil {
			panic(err)
		}
		log.Printf("User \"%s\" logged in succesfully", userName)
		http.Redirect(w, r, "/userarea", http.StatusFound)
	} else {
		log.Printf("User \"%s\" log in attemp failed", userName)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func (c Controller) UserAreaHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/userarea.html")
	cookieUsername, err := r.Cookie("QuizUser")
	if err != nil {
		log.Println("Wrong Cookie :", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	cookieToken, err := r.Cookie("QuizUserToken")
	if err != nil {
		log.Println("Wrong Cookie :", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	result, err := checkUserCookie(c, cookieUsername.Value, cookieToken.Value)
	if err != nil {
		log.Println("Wrong Cookie :", err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if result == true {
		tasks, _ := database.GetTasks(c.Db)
		err = t.Execute(w, tasks)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error parse template :", err)
			return
		}
	} else {
		log.Println("Wrong Cookie :", err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
