package main

import(
  "bytes"
  "database/sql"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestRequestSearchNickname(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  librarian.add(Person{
    Nickname: "Cu Tuấn",
    FirstName: sql.NullString{String: "Tuấn", Valid: true,},
    Description: sql.NullString{String: "khá bảnh ;))", Valid: true,},
  })
  requestPtr := httptest.NewRequest( http.MethodGet,
                                     "/search",
                                     bytes.NewBuffer( []byte(`{"nickname": "Cu Tuấn"}`) ) )
  people, err := search(requestPtr, &librarian)
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if len(people) != 1 {
    t.Errorf("expected 1 got %v", len(people))
  }
  if people[0].Nickname != "Cu Tuấn" {
    t.Errorf("expected \"Cu Tuấn\" got %v", people[0].Nickname)
  }
  if people[0].FirstName.String != "Tuấn" {
    t.Errorf("expected Tuấn got %v", people[0].FirstName.String)
  }
  if people[0].Description.String != "khá bảnh ;))" {
    t.Errorf("expected \"khá bảnh ;))\" (without quotes) got %v", people[0].Description.String)
  }
}

func TestRequestSearchNicknameNotExist(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  requestPtr := httptest.NewRequest( http.MethodGet,
                                     "/search",
                                     bytes.NewBuffer( []byte(`{"nickname": "凱哥"}`) ) )
  people, err := search(requestPtr, &librarian)
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if len(people) != 0 {
    t.Errorf("expected 0 got %v", len(people))
  }
}
