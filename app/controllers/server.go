package controllers

import (
  "fmt"
  "html/template"
  "net/http"
  "../../config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
  var files []string
  for _, file := range filenames {
    files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
  }
  templates := template.Must(template.ParseFiles(files...))
  templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
  files := http.FileServer(http.Dir(config.Config.Static))

  // staticというフォルダは存在しないためStripPrefixで検索対象から除外する
  http.Handle("/static/", http.StripPrefix("/static/", files))
  http.HandleFunc("/", top)
  return http.ListenAndServe(":"+config.Config.Port, nil)
}
