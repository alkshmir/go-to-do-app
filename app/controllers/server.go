package controllers

import (
	"fmt"
	"go-to-do-app/config"
	"html/template"
	"net/http"
	"strconv"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...)) //キャッシュ
	templates.ExecuteTemplate(w, "layout", data)              //実行するテンプレートを明示的に指定
}
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top) //url登録
	return http.ListenAndServe(":"+strconv.Itoa(config.Config.Port), nil)
}
