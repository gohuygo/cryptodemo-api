package main

import (
  "fmt"
  "log"
  "net/http"
  "os"

  "encoding/json"
  "./models"
  "./controllers"

  "github.com/dgrijalva/jwt-go"
  // "github.com/gorilla/context"
  "github.com/gorilla/mux"
  // "github.com/mitchellh/mapstructure"
)

type JwtToken struct {
  Token string `json:"token"`
}

type Exception struct {
  Message string `json:"message"`
}

func setupResponse(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  w.WriteHeader(http.StatusCreated)
}

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Create Token Endpoint")
  var user models.User

  // TODO: Validate body has name/pw
  _ = json.NewDecoder(req.Body).Decode(&user)

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "sub":  user.Name,
    "iss": "cryptodemo",
    //"exp": Time
  })

  tokenString, err := token.SignedString([]byte("testsecret"))
  if err != nil { fmt.Println(err) }
  setupResponse(w, req)

  json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
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

  router.HandleFunc("/authenticate",  CreateTokenEndpoint).Methods("POST")
  router.HandleFunc("/protected",  ProtectedEndpoint).Methods("GET")
  router.HandleFunc("/", controller.IndexEndpoint)

  log.Fatal(http.ListenAndServe(":"+port, router))
}

