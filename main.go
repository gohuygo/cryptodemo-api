package main

import (
  "fmt"
  "log"
  "net/http"
  "os"

  "encoding/json"
  "./controllers"

  "github.com/dgrijalva/jwt-go"
  // "github.com/gorilla/context"
  "github.com/gorilla/mux"
  // "github.com/mitchellh/mapstructure"
)



type Exception struct {
  Message string `json:"message"`
}


func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Protected Endpoint!!")
  params := req.URL.Query()

  token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("There was an error")
    }
    return []byte("testsecret"), nil
  })

  fmt.Println(params)

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      // var user User
      // mapstructure.Decode(claims, &user)
      json.NewEncoder(w).Encode(claims)
  } else {
      json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
  }
}


func main() {
  // Set production port (Heroku)
  port := os.Getenv("PORT")

  // Set development port
  if port  == "" {
    port = "8080"
  }

  router := mux.NewRouter()
  fmt.Println("Starting application...")

  homeController := controllers.NewHomeController()
  authenticationController := controllers.NewAuthenticationController()

  router.HandleFunc("/authenticate",  authenticationController.CreateTokenEndpoint).Methods("POST")
  router.HandleFunc("/protected",  ProtectedEndpoint).Methods("GET")
  router.HandleFunc("/", homeController.IndexEndpoint)

  log.Fatal(http.ListenAndServe(":"+port, router))
}

