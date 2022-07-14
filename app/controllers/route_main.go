package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		// ログインしていない
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
	/*
		t, err := template.ParseFiles("app/views/templates/top.html")
		if err != nil {
			log.Fatalln(err)
		}
		t.Execute(w, "Hello")
	*/
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
