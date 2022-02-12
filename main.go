package main

import(
  "encoding/json"
  "net/http"
  "errors"
)

func search(requestPtr *http.Request, librarianPtr *Librarian) (*Person, error) {
  var person Person
  err := json.NewDecoder(requestPtr.Body).Decode(&person)
  if err != nil {
    return nil, err
  }
  people, err := librarianPtr.search(person.Nickname)
  if len(people) == 0 {
    return nil, errors.New("search 凱哥: unknown nickname")
  } else {
    return &people[0], err
  }
}

func main() {
}
