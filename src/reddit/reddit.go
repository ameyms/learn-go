package reddit

import (
 "net/http"
  "log"
  "encoding/json"
)

// Wraps the complete response of the reddit API
type Response struct {
  Data struct{
    Children []struct {
      Data Item
    }
  }
}

// Each item from the reddit API response
type Item struct {
  Title string
  URL string
}

func Read(topic string) (*Response, error) {
  r, err := http.Get("http://reddit.com/r/" + topic + ".json")
  if err != nil {
    log.Fatal(err)
    return nil, err
  }
  if r.StatusCode != http.StatusOK {
    log.Fatal(r.Status)
    return nil, err
  }

  resp := new(Response)
  err = json.NewDecoder(r.Body).Decode(resp)

  return resp, err
}
