package reddit

import (
 "net/http"
  "log"
  "fmt"
  "encoding/json"
  "errors"
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
  Comments int `json:"num_comments"`
}

func (i Item) String() string {
  return fmt.Sprintf("%s (%d)\n%s", i.Title, i.Comments, i.URL)
}

func Read(topic string) (*Response, error) {
  r, err := http.Get("http://reddit.com/r/" + topic + ".json")
  if err != nil {
    log.Fatal(err)
    return nil, err
  }
  //Cleanup our http response object
  defer r.Body.Close()

  if r.StatusCode != http.StatusOK {
    log.Fatal(r.Status)
    return nil, errors.New(r.Status)
  }

  resp := new(Response)
  err = json.NewDecoder(r.Body).Decode(resp)

  return resp, err
}
