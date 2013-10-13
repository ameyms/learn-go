package main

import (
  "time"
  "net/http"
  "io/ioutil"
  "io"
  "fmt"
)

// Holds information for website
type WebSite struct {
  Name string
  URL string
}

func (w *WebSite) String() string {
  return fmt.Sprintf("%s (%s)", w.Name, w.URL)
}


// Holds page statistics
type PageStats struct {

  Size int64
  LoadTime time.Duration
  Err error
  Page *WebSite
}

// Implement Stringer
func (ps PageStats) String() string {

  return fmt.Sprintf("%s took %s and is %d bytes", ps.Page, ps.LoadTime, ps.Size)
}


// HTTP GETs a particular website and publishes stats to channel
func GetStats(url *WebSite, c chan PageStats) {

  start := time.Now()
  response, err := http.Get(url.URL)

  if err != nil {
    c <- PageStats{Page:url, Size:0, LoadTime:time.Since(start), Err:err}
  }

  n, _ := io.Copy(ioutil.Discard, response.Body)
  defer response.Body.Close()

  c <- PageStats{Page:url, LoadTime:time.Since(start), Size:n}
}

// TODO: Investigate why changing method signature to *WebSite doesn`t work
func doGetStats(w WebSite, c chan PageStats) {
    go GetStats(&w, c)
}

func main() {

  var websites = []WebSite  {WebSite{"Google", "http://google.com"}, WebSite{"Facebook", "http://facebook.com"}}
  start := time.Now()

  c := make(chan PageStats, 5)

  for _, w := range websites {
    doGetStats(w, c)
 }
 for i := 0; i < len(websites); i++ {
  fmt.Printf("%s\n", <-c)
 }
 fmt.Printf("Total Time taken: %s\n", time.Since(start))
}
