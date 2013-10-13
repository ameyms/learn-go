package main

import "fmt"

type Language int

const (
  JavaScript Language = iota
  Go
  Java
  Python
)

var lang = []string {"Javascript", "GoLang", "Java", "Python"}

func (l Language) String() string {
  return "Hack using " + lang[l]
}

func main() {
  fmt.Printf("Are you a programmer? %s\n", Go)
}
