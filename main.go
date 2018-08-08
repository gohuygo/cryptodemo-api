package main

import(
  cmc "github.com/coincircle/go-coinmarketcap"
  "fmt"
  "net/http"
  "os"
  "time"
  "strconv"
)

func main() {
  port := os.Getenv("PORT") 

  if port  == "" {
    port = "8080"
  }
  
  http.HandleFunc("/", index)
  
  http.ListenAndServe(":" + port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
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

  // fmt.Println(graph.PriceBTC)
  fmt.Fprintln(w, graph.PriceBTC)

}
