package main

import(
  "bytes"
  "database/sql"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestRequestAddPerson(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  requestPtr := httptest.NewRequest( http.MethodPost,
                                     "/add",
                                     bytes.NewBuffer( []byte(`{"nickname": "啊威", "firstName": {"String": "威", "Valid": true}, "middleName": {"String": "", "Valid": false}, "lastName": {"String": "劉", "Valid": true}, "phoneCountry": {"String": "", "Valid": false}, "phoneArea": {"String": "", "Valid": false}, "phoneNumber": {"String": "999888666", "Valid": true}, "email": {"String": "something@google.com", "Valid": true}, "birthdate": {"String": "", "Valid": false}, "description": {"String": "", "Valid": false}}`) ) )
  err := add(requestPtr, &librarian)
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  people, err := librarian.search("啊威")
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if len(people) != 1 {
    t.Errorf("expected 1 got %v", len(people))
  }
  if people[0].Nickname != "啊威" {
    t.Errorf("expected 啊威 got %v", people[0].Nickname)
  }
  if people[0].FirstName.String != "威" {
    t.Errorf("expected 威 got %v", people[0].FirstName.String)
  }
  if !people[0].FirstName.Valid {
    t.Errorf("expected true got %v", people[0].FirstName.Valid)
  }
  if people[0].MiddleName.String != "" {
    t.Errorf("expected empty string got %v", people[0].MiddleName.String)
  }
  if people[0].MiddleName.Valid {
    t.Errorf("expected false got %v", people[0].MiddleName.Valid)
  }
  if people[0].LastName.String != "劉" {
    t.Errorf("expected 劉 got %v", people[0].LastName.String)
  }
  if !people[0].LastName.Valid {
    t.Errorf("expected true got %v", people[0].LastName.Valid)
  }
  if people[0].PhoneCountry.String != "" {
    t.Errorf("expected empty string got %v", people[0].PhoneCountry.String)
  }
  if people[0].PhoneCountry.Valid {
    t.Errorf("expected false got %v", people[0].PhoneCountry.Valid)
  }
  if people[0].PhoneArea.String != "" {
    t.Errorf("expected empty string got %v", people[0].PhoneArea.String)
  }
  if people[0].PhoneArea.Valid {
    t.Errorf("expected false got %v", people[0].PhoneArea.Valid)
  }
  if people[0].PhoneNumber.String != "999888666" {
    t.Errorf("expected 999888666 got %v", people[0].PhoneNumber.String)
  }
  if !people[0].PhoneNumber.Valid {
    t.Errorf("expected true got %v", people[0].PhoneNumber.Valid)
  }
  if people[0].Email.String != "something@google.com" {
    t.Errorf("expected something@google.com got %v", people[0].Email.String)
  }
  if !people[0].PhoneNumber.Valid {
    t.Errorf("expected true got %v", people[0].PhoneNumber.Valid)
  }
  if people[0].Birthdate.String != "" {
    t.Errorf("expected empty string got %v", people[0].Birthdate.String)
  }
  if people[0].Birthdate.Valid {
    t.Errorf("expected false got %v", people[0].Birthdate.Valid)
  }
  if people[0].Description.String != "" {
    t.Errorf("expected empty string got %v", people[0].Description.String)
  }
  if people[0].Description.Valid {
    t.Errorf("expected false got %v", people[0].Description.Valid)
  }
}

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
