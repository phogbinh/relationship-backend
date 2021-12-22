package main

import(
  "encoding/json"
  "net/http"
)

func search(requestPtr *http.Request, librarianPtr *Librarian) (*Person, error) {
  var person Person
  err := json.NewDecoder(requestPtr.Body).Decode(&person)
  if err != nil {
    return nil, err
  }
  return librarianPtr.search(person.Nickname)
}

func main() {
}
