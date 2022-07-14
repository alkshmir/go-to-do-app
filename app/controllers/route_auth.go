package controllers

import (
	"go-to-do-app/app/models"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			// ログインしていない
			generateHTML(w, "Hello", "layout", "signup")
		} else {
			http.Redirect(w, r, "/todos", http.StatusFound)
		}
		//generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Fatalln(err)
		}

		http.Redirect(w, r, "/", http.StatusFound) //トップにリダイレクトさせる
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		// ログインしていない
		generateHTML(w, "Hello", "layout", "login")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Fatalln(err)
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	if user.PassWord == models.Hash(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Fatalln(err)
		}

		//クッキー作成
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusBadRequest)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
