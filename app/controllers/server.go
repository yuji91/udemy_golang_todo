package controllers

import (
  "net/http"
  "../../config"
)

func StartMainServer() error {
  files := http.FileServer(http.Dir(config.Config.Static))

  // staticというフォルダは存在しないためStripPrefixで検索対象から除外する
  http.Handle("/static/", http.StripPrefix("/static/", files))
  http.HandleFunc("/", top)
  return http.ListenAndServe(":"+config.Config.Port, nil)
}
