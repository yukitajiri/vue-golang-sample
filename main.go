package main

import (
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "net/http"
  "log"
)

func main () {
  allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8081"})
  allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
  allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})


  r := mux.NewRouter()

  r.HandleFunc("/public", public)
  r.HandleFunc("/private", pravate)

  log.Fatal(http.ListenAndServe(":8081", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}

func public (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hello public!\n"))
}

func pravate (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hello private!\n"))
}
