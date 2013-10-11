package reddit

import (
  "testing"
  "fmt"
)

func TestExisting(t *testing.T) {
  resp, err := Read("javascript")
  if err != nil {
    t.Error(err)

  } else {
    for _, child := range resp.Data.Children {
      fmt.Println(child.Data)
    }
  }
}



