package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  html_file, err := ioutil.ReadFile("test.html")
  check(err)
  fmt.Fprintln(w, string(html_file))
}


func serve(lab_status map[string][]map[string]int) {
  http.HandleFunc("/initial", handler)
  http.ListenAndServe(":8080",nil)
}

