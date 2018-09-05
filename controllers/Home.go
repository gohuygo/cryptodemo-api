package controller

import (
  "fmt"
  "encoding/json"

  "net/http"
  "time"
  "strconv"
  cmc "github.com/coincircle/go-coinmarketcap"

)

func IndexEndpoint(w http.ResponseWriter, r *http.Request) {
  threeMonths := int64(60 * 60 * 24 * 90)
  now := time.Now()
  secs := now.Unix()
  start := secs - threeMonths
  end := secs

  fmt.Println("Time is " + strconv.FormatInt(end, 10))

  graph, _ := cmc.TickerGraph(&cmc.TickerGraphOptions{
    Start: start,
    End: end,
    Symbol: "ETH",
  })

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)

  json.NewEncoder(w).Encode(graph)
}
