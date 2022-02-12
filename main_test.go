package main

import(
  "database/sql"
  "testing"
  "net/http"
  "net/http/httptest"
  "bytes"
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
  personPtr, err := search(requestPtr, &librarian)
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if personPtr.Nickname != "Cu Tuấn" {
    t.Errorf("expected \"Cu Tuấn\" got %v", personPtr.Nickname)
  }
  if personPtr.FirstName.String != "Tuấn" {
    t.Errorf("expected Tuấn got %v", personPtr.FirstName.String)
  }
  if personPtr.Description.String != "khá bảnh ;))" {
    t.Errorf("expected \"khá bảnh ;))\" (without quotes) got %v", personPtr.Description.String)
  }
}

func TestRequestSearchNicknameNotExist(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  requestPtr := httptest.NewRequest( http.MethodGet,
                                     "/search",
                                     bytes.NewBuffer( []byte(`{"nickname": "凱哥"}`) ) )
  personPtr, err := search(requestPtr, &librarian)
  if personPtr != nil {
    t.Errorf("expected nil got %v", personPtr)
  }
  if err.Error() != "search 凱哥: unknown nickname" {
    t.Errorf("expected \"search 凱哥: unknown nickname\" (without quotes) got %v", err.Error())
  }
}
