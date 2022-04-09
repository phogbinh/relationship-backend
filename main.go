package main

import(
  "database/sql"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"

  "github.com/go-sql-driver/mysql"
)

var librarian Librarian

func add(requestPtr *http.Request, librarianPtr *Librarian) (*Person, error) {
  var person Person
  err := json.NewDecoder(requestPtr.Body).Decode(&person)
  if err != nil {
    return nil, err
  }
  return librarianPtr.add(person)
}

func handleRequestAddPerson(responseWriter http.ResponseWriter, requestPtr *http.Request) {
  personPtr, err := add(requestPtr, &librarian)
  if err != nil {
    responseWriter.WriteHeader(http.StatusInternalServerError)
    fmt.Println(err)
    return
  }
  responseWriter.WriteHeader(http.StatusCreated)
  err = json.NewEncoder(responseWriter).Encode(personPtr)
  if err != nil { // TODO test
    responseWriter.WriteHeader(http.StatusInternalServerError)
    fmt.Println(err)
    return
  }
}

func search(requestPtr *http.Request, librarianPtr *Librarian) ([]Person, error) {
  var person Person
  err := json.NewDecoder(requestPtr.Body).Decode(&person) // TODO detect incorrect json
  if err != nil {
    return nil, err
  }
  return librarianPtr.search(person.Nickname)
}

func handleRequestSearchNickname(responseWriter http.ResponseWriter, requestPtr *http.Request) {
  people, err := search(requestPtr, &librarian)
  if err != nil {
    responseWriter.WriteHeader(http.StatusInternalServerError)
    fmt.Println(err)
    return
  }
  err = json.NewEncoder(responseWriter).Encode(ResponseBody{People: people,})
  if err != nil { // TODO test
    responseWriter.WriteHeader(http.StatusInternalServerError)
    fmt.Println(err)
    return
  }
}

func update(requestPtr *http.Request, librarianPtr *Librarian) (*Person, error) {
  var person Person
  err := json.NewDecoder(requestPtr.Body).Decode(&person)
  if err != nil {
    return nil, err
  }
  return librarianPtr.update(person)
}

func main() {
  config := mysql.Config{
    User: os.Getenv("RELATIONSHIP_BACKEND_DATABASE_USER"),
    Passwd: os.Getenv("RELATIONSHIP_BACKEND_DATABASE_PASSWORD"),
    Net: "tcp",
    Addr: os.Getenv("RELATIONSHIP_BACKEND_DATABASE_ADDRESS"),
    DBName: os.Getenv("RELATIONSHIP_BACKEND_DATABASE_NAME"),
  }
  sqlDatabasePtr, err := sql.Open("mysql", config.FormatDSN())
  if err != nil {
    log.Fatal(err)
  }
  database := Database{
    SqlDatabasePtr: sqlDatabasePtr,
  }
  fmt.Println("database " + os.Getenv("RELATIONSHIP_BACKEND_DATABASE_NAME") + " connected")
  librarian = Librarian{
    DatabasePtr: &database,
  }
  http.HandleFunc("/search", handleRequestSearchNickname)
  http.HandleFunc("/add", handleRequestAddPerson)
  log.Fatal(http.ListenAndServe(os.Getenv("RELATIONSHIP_BACKEND_PORT"), nil))
}
