package main

import (
  "fmt"
  "log"
  "net/http"
  "os"

  // Docker issue with referencing local files?
  // "./controllers"
  // "./middleware"
  "github.com/gohuygo/cryptodemo-api/controllers"
  "github.com/gohuygo/cryptodemo-api/middleware"
  "github.com/gorilla/mux"

  "database/sql"
  _ "github.com/lib/pq"
)


func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Protected Endpoint!!")
  w.Header().Set("Content-Type", "application/json")

  // json.NewEncoder(w).Encode([]byte("Accessed Protected Endpoint!"))
  w.Write([]byte("Accessed Protected Endpoint!"))
}

func main() {
  // Set production port (Heroku)
  port := os.Getenv("PORT")

  // Set development port
  if port  == "" {
    port = "80"
  }

  db, err := sql.Open("postgres", "postgres://test:password@localhost/cryptodemo?sslmode=disable")

  router := mux.NewRouter()
  fmt.Println("Starting application...")

  tokenValidator := middleware.NewTokenValidator()

  homeController := controllers.NewHomeController()
  authenticationController := controllers.NewAuthenticationController()

  router.HandleFunc("/authenticate",  authenticationController.CreateTokenEndpoint).Methods("POST")
  router.HandleFunc("/protected",  tokenValidator.Validate(ProtectedEndpoint)).Methods("GET")
  router.HandleFunc("/", homeController.IndexEndpoint)

  log.Fatal(http.ListenAndServe(":"+port, router))
}

