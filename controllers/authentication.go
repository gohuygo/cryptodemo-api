package controllers

import(
  "fmt"
  "net/http"
  "encoding/json"

   // "../models"
  "github.com/gohuygo/cryptodemo-api/models"
  "github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
  Token string `json:"token"`
}

type AuthenticationController struct{
  // TODO: Connection to DB
}

func NewAuthenticationController() *AuthenticationController {
  return &AuthenticationController{}
}

func (ac AuthenticationController) CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
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

func setupResponse(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  w.WriteHeader(http.StatusCreated)
}
