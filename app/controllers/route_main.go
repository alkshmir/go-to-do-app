package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "Hello", "layout", "top")
	/*
		t, err := template.ParseFiles("app/views/templates/top.html")
		if err != nil {
			log.Fatalln(err)
		}
		t.Execute(w, "Hello")
	*/
}
