package main

import(
  "fmt"
  "log"
  "net/http"
  // "os"
  // "time"
  // "strconv"
  // "encoding/json"

  // cmc "github.com/coincircle/go-coinmarketcap"

  // "github.com/dgrijalva/jwt-go"
  // "github.com/gorilla/context"
  "github.com/gorilla/mux"
  // "github.com/mitchellh/mapstructure"
)

type User struct {
  Name string
  Password string
}

type JwtToken struct {
  Token string
}

func CreateTokenEndpoint(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Create Token Endpoint")
}

func ProtectedEndpoint(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Protected Endpoint")
}

func main() {
  router := mux.NewRouter()
  fmt.Println("Starting application...")

  router.HandleFunc("/authenticate",  CreateTokenEndpoint).Methods("POST")
  router.HandleFunc("/protected",  ProtectedEndpoint).Methods("GET")

  log.Fatal(http.ListenAndServe(":8080", router))
}

// func main() {
//   port := os.Getenv("PORT")

//   if port  == "" {
//     port = "8080"
//   }

//   http.HandleFunc("/", index)

//   http.ListenAndServe(":" + port, nil)
// }

// func index(w http.ResponseWriter, r *http.Request) {
//   threeMonths := int64(60 * 60 * 24 * 90)
//   now := time.Now()
//   secs := now.Unix()
//   start := secs - threeMonths
//   end := secs

//   fmt.Println("Time is " + strconv.FormatInt(end, 10))

//   graph, _ := cmc.TickerGraph(&cmc.TickerGraphOptions{
//     Start: start,
//     End: end,
//     Symbol: "ETH",
//   })

//   w.Header().Set("Content-Type", "application/json")
//   w.WriteHeader(http.StatusCreated)

//   json.NewEncoder(w).Encode(graph)
// }
