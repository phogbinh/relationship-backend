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
  personPtr, err := add(requestPtr, &librarian)
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if personPtr.Nickname != "啊威" {
    t.Errorf("expected 啊威 got %v", personPtr.Nickname)
  }
  if personPtr.FirstName.String != "威" {
    t.Errorf("expected 威 got %v", personPtr.FirstName.String)
  }
  if !personPtr.FirstName.Valid {
    t.Errorf("expected true got %v", personPtr.FirstName.Valid)
  }
  if personPtr.MiddleName.String != "" {
    t.Errorf("expected empty string got %v", personPtr.MiddleName.String)
  }
  if personPtr.MiddleName.Valid {
    t.Errorf("expected false got %v", personPtr.MiddleName.Valid)
  }
  if personPtr.LastName.String != "劉" {
    t.Errorf("expected 劉 got %v", personPtr.LastName.String)
  }
  if !personPtr.LastName.Valid {
    t.Errorf("expected true got %v", personPtr.LastName.Valid)
  }
  if personPtr.PhoneCountry.String != "" {
    t.Errorf("expected empty string got %v", personPtr.PhoneCountry.String)
  }
  if personPtr.PhoneCountry.Valid {
    t.Errorf("expected false got %v", personPtr.PhoneCountry.Valid)
  }
  if personPtr.PhoneArea.String != "" {
    t.Errorf("expected empty string got %v", personPtr.PhoneArea.String)
  }
  if personPtr.PhoneArea.Valid {
    t.Errorf("expected false got %v", personPtr.PhoneArea.Valid)
  }
  if personPtr.PhoneNumber.String != "999888666" {
    t.Errorf("expected 999888666 got %v", personPtr.PhoneNumber.String)
  }
  if !personPtr.PhoneNumber.Valid {
    t.Errorf("expected true got %v", personPtr.PhoneNumber.Valid)
  }
  if personPtr.Email.String != "something@google.com" {
    t.Errorf("expected something@google.com got %v", personPtr.Email.String)
  }
  if !personPtr.PhoneNumber.Valid {
    t.Errorf("expected true got %v", personPtr.PhoneNumber.Valid)
  }
  if personPtr.Birthdate.String != "" {
    t.Errorf("expected empty string got %v", personPtr.Birthdate.String)
  }
  if personPtr.Birthdate.Valid {
    t.Errorf("expected false got %v", personPtr.Birthdate.Valid)
  }
  if personPtr.Description.String != "" {
    t.Errorf("expected empty string got %v", personPtr.Description.String)
  }
  if personPtr.Description.Valid {
    t.Errorf("expected false got %v", personPtr.Description.Valid)
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

func TestRequestUpdateId(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  librarian.add(Person{
    Id: 3,
    Nickname: "Jen",
  })
  requestPtr := httptest.NewRequest( http.MethodPut,
                                     "/update",
                                     bytes.NewBuffer( []byte(`{"id": 3, "nickname": "哥", "middleName": {"String": "@%\\", "Valid": true}}`) ) )
  personPtr, err := update(requestPtr, &librarian)
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if personPtr.Id != 3 {
    t.Errorf("expected 3 got %v", personPtr.Id)
  }
  if personPtr.Nickname != "哥" {
    t.Errorf("expected 哥 got %v", personPtr.Nickname)
  }
  if personPtr.MiddleName.String != "@%\\" {
    t.Errorf("expected @%%\\ got %v", personPtr.MiddleName.String)
  }
}

func TestRequestUpdateIdNotExist(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  requestPtr := httptest.NewRequest( http.MethodPut,
                                     "/update",
                                     bytes.NewBuffer( []byte(`{"id": 1, "nickname": "something"}`) ) )
  personPtr, err := update(requestPtr, &librarian)
  if err == nil {
    t.Errorf("expected error got nil")
  }
  if personPtr != nil {
    t.Errorf("expected nil got %v", personPtr)
  }
}
