package main

import(
  "testing"
)

func TestAddFullPerson(t *testing.T) {
  db := MockDatabase{}
  librarian := Librarian{
    DatabasePtr: &db,
  }
  err := librarian.add(Person{
    Nickname: "Bullshit",
    FirstName: "Lam",
    MiddleName: "Nha",
    LastName: "Tranh",
    PhoneCountry: "84",
    PhoneArea: "2",
    PhoneNumber: "111222333",
    Email: "nhatrang@gmail.com",
    Birthdate: "19970725",
    Description: "jerk",
  })
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  rowsPtr, err := db.Query("SELECT * FROM person WHERE nickname = \"Bullshit\"")
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  defer rowsPtr.Close()
  var people []Person
  for rowsPtr.Next() {
    var id int
    var person Person
    err := rowsPtr.Scan(&id,
                        &person.Nickname,
                        &person.FirstName,
                        &person.MiddleName,
                        &person.LastName,
                        &person.PhoneCountry,
                        &person.PhoneArea,
                        &person.PhoneNumber,
                        &person.Email,
                        &person.Birthdate,
                        &person.Description)
    if err != nil {
      t.Errorf("expected error to be nil got %v", err)
    }
    people = append(people, person)
  }
  err = rowsPtr.Err()
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if len(people) != 1 {
    t.Errorf("expected 1 got %v", len(people))
  }
  person := people[0]
  if person.Nickname != "Bullshit" {
    t.Errorf("expected Bullshit got %v", person.Nickname)
  }
  if person.FirstName != "Lam" {
    t.Errorf("expected Lam got %v", person.MiddleName)
  }
  if person.MiddleName != "Nha" {
    t.Errorf("expected Nha got %v", person.MiddleName)
  }
  if person.LastName != "Tranh" {
    t.Errorf("expected Tranh got %v", person.LastName)
  }
  if person.PhoneCountry != "84" {
    t.Errorf("expected 84 got %v", person.PhoneCountry)
  }
  if person.PhoneArea != "2" {
    t.Errorf("expected 2 got %v", person.PhoneArea)
  }
  if person.PhoneNumber != "111222333" {
    t.Errorf("expected 111222333 got %v", person.PhoneNumber)
  }
  if person.Email != "nhatrang@gmail.com" {
    t.Errorf("expected nhatrang@gmail.com got %v", person.Email)
  }
  if person.Birthdate != "19970725" {
    t.Errorf("expected 19970725 got %v", person.Birthdate)
  }
  if person.Description != "jerk" {
    t.Errorf("expected jerk got %v", person.Description)
  }
}

func TestAddPartialPerson(t *testing.T) {
  db := MockDatabase{}
  librarian := Librarian{
    DatabasePtr: &db,
  }
  err := librarian.add(Person{
    Nickname: "Johnny",
    FirstName: "John",
    Description: "he seems like a nice guy",
  })
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  rowsPtr, err := db.Query("SELECT * FROM person WHERE nickname = \"Johnny\"")
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  defer rowsPtr.Close()
  var people []Person
  for rowsPtr.Next() {
    var id int
    var person Person
    err := rowsPtr.Scan(&id,
                        &person.Nickname,
                        &person.FirstName,
                        &person.MiddleName,
                        &person.LastName,
                        &person.PhoneCountry,
                        &person.PhoneArea,
                        &person.PhoneNumber,
                        &person.Email,
                        &person.Birthdate,
                        &person.Description)
    if err != nil {
      t.Errorf("expected error to be nil got %v", err)
    }
    people = append(people, person)
  }
  err = rowsPtr.Err()
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if len(people) != 1 {
    t.Errorf("expected 1 got %v", len(people))
  }
  person := people[0]
  if person.Nickname != "Johnny" {
    t.Errorf("expected Johnny got %v", person.Nickname)
  }
  if person.FirstName != "John" {
    t.Errorf("expected John got %v", person.MiddleName)
  }
  if person.MiddleName != "" {
    t.Errorf("expected empty string got %v", person.MiddleName)
  }
  if person.LastName != "" {
    t.Errorf("expected empty string got %v", person.LastName)
  }
  if person.PhoneCountry != "" {
    t.Errorf("expected empty string got %v", person.PhoneCountry)
  }
  if person.PhoneArea != "" {
    t.Errorf("expected empty string got %v", person.PhoneArea)
  }
  if person.PhoneNumber != "" {
    t.Errorf("expected empty string got %v", person.PhoneNumber)
  }
  if person.Email != "" {
    t.Errorf("expected empty string got %v", person.Email)
  }
  if person.Birthdate != "" {
    t.Errorf("expected empty string got %v", person.Birthdate)
  }
  if person.Description != "he seems like a nice guy" {
    t.Errorf("expected \"he seems like a nice guy\" (without quotes) got %v", person.Description)
  }
}

func test_search_nickname(t *testing.T) {
  librarian := Librarian{}
  person := Person{
    Nickname: "Bullshit",
    FirstName: "Lam",
    MiddleName: "Nha",
    LastName: "Tranh",
    PhoneCountry: "84",
    PhoneArea: "2",
    PhoneNumber: "111222333",
    Email: "nhatrang@gmail.com",
    Birthdate: "19970725",
    Description: "jerk",
  }
  err := librarian.add(person)
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  person, err = librarian.search("Bullshit")
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if person.Nickname != "Bullshit" {
    t.Errorf("expected Bullshit got %v", person.Nickname)
  }
  if person.FirstName != "Lam" {
    t.Errorf("expected Lam got %v", person.MiddleName)
  }
  if person.MiddleName != "Nha" {
    t.Errorf("expected Nha got %v", person.MiddleName)
  }
  if person.LastName != "Tranh" {
    t.Errorf("expected Tranh got %v", person.LastName)
  }
  if person.PhoneCountry != "84" {
    t.Errorf("expected 84 got %v", person.PhoneCountry)
  }
  if person.PhoneArea != "2" {
    t.Errorf("expected 2 got %v", person.PhoneArea)
  }
  if person.PhoneNumber != "111222333" {
    t.Errorf("expected 111222333 got %v", person.PhoneNumber)
  }
  if person.Email != "nhatrang@gmail.com" {
    t.Errorf("expected nhatrang@gmail.com got %v", person.Email)
  }
  if person.Birthdate != "19970725" {
    t.Errorf("expected 19970725 got %v", person.Birthdate)
  }
  if person.Description != "jerk" {
    t.Errorf("expected jerk got %v", person.Description)
  }
}
