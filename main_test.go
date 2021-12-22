package main

import(
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
    FirstName: "Tuấn",
    Description: "khá bảnh ;))",
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
  if personPtr.FirstName != "Tuấn" {
    t.Errorf("expected Tuấn got %v", personPtr.FirstName)
  }
  if personPtr.MiddleName != "" {
    t.Errorf("expected empty string got %v", personPtr.MiddleName)
  }
  if personPtr.LastName != "" {
    t.Errorf("expected empty string got %v", personPtr.LastName)
  }
  if personPtr.PhoneCountry != "" {
    t.Errorf("expected empty string got %v", personPtr.PhoneCountry)
  }
  if personPtr.PhoneArea != "" {
    t.Errorf("expected empty string got %v", personPtr.PhoneArea)
  }
  if personPtr.PhoneNumber != "" {
    t.Errorf("expected empty string got %v", personPtr.PhoneNumber)
  }
  if personPtr.Email != "" {
    t.Errorf("expected empty string got %v", personPtr.Email)
  }
  if personPtr.Birthdate != "" {
    t.Errorf("expected empty string got %v", personPtr.Birthdate)
  }
  if personPtr.Description != "khá bảnh ;))" {
    t.Errorf("expected \"khá bảnh ;))\" (without quotes) got %v", personPtr.Description)
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
