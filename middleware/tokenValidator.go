package middleware

import (
  "fmt"
  "net/http"
  "strings"
  "encoding/json"

  "github.com/dgrijalva/jwt-go"
  "github.com/gorilla/context"
)

type Exception struct {
  Message string `json:"message"`
}

type TokenValidator struct{}

func NewTokenValidator() *TokenValidator {
  return &TokenValidator{}
}

func (tv TokenValidator) Validate(endpoint http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    authorizationHeader := req.Header.Get("authorization")
    if authorizationHeader != "" {
      bearerToken := strings.Split(authorizationHeader, " ")
      if len(bearerToken) == 2 {
        token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
          _, ok := token.Method.(*jwt.SigningMethodHMAC)

          if !ok {
            return nil, fmt.Errorf("There was an error")
          }

          return []byte("testsecret"), nil
        })
        if error != nil {
          json.NewEncoder(w).Encode(Exception{Message: error.Error()})
          return
        }

        if token.Valid {
          context.Set(req, "decoded", token.Claims)
          endpoint(w, req) //if valid, call the endpoint
        } else {
          json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
        }
      }
    } else {
      json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
    }
  })
}
