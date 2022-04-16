package main

import(
  "database/sql"
  "testing"
)

func TestAddFullPerson(t *testing.T) {
  db := MockDatabase{}
  librarian := Librarian{
    DatabasePtr: &db,
  }
  personPtr, err := librarian.add(Person{
    Nickname: "Bullshit",
    FirstName: sql.NullString{String: "Lam", Valid: true,},
    MiddleName: sql.NullString{String: "Nha", Valid: true,},
    LastName: sql.NullString{String: "Tranh", Valid: true,},
    PhoneCountry: sql.NullString{String: "84", Valid: true,},
    PhoneArea: sql.NullString{String: "2", Valid: true,},
    PhoneNumber: sql.NullString{String: "111222333", Valid: true,},
    Email: sql.NullString{String: "nhatrang@gmail.com", Valid: true,},
    Birthdate: sql.NullString{String: "19970725", Valid: true,},
    Description: sql.NullString{String: "jerk", Valid: true,},
  })
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if personPtr.Nickname != "Bullshit" {
    t.Errorf("expected Bullshit got %v", personPtr.Nickname)
  }
  if personPtr.FirstName.String != "Lam" {
    t.Errorf("expected Lam got %v", personPtr.FirstName.String)
  }
  if personPtr.MiddleName.String != "Nha" {
    t.Errorf("expected Nha got %v", personPtr.MiddleName.String)
  }
  if personPtr.LastName.String != "Tranh" {
    t.Errorf("expected Tranh got %v", personPtr.LastName.String)
  }
  if personPtr.PhoneCountry.String != "84" {
    t.Errorf("expected 84 got %v", personPtr.PhoneCountry.String)
  }
  if personPtr.PhoneArea.String != "2" {
    t.Errorf("expected 2 got %v", personPtr.PhoneArea.String)
  }
  if personPtr.PhoneNumber.String != "111222333" {
    t.Errorf("expected 111222333 got %v", personPtr.PhoneNumber.String)
  }
  if personPtr.Email.String != "nhatrang@gmail.com" {
    t.Errorf("expected nhatrang@gmail.com got %v", personPtr.Email.String)
  }
  if personPtr.Birthdate.String != "19970725" {
    t.Errorf("expected 19970725 got %v", personPtr.Birthdate.String)
  }
  if personPtr.Description.String != "jerk" {
    t.Errorf("expected jerk got %v", personPtr.Description.String)
  }
}

func TestAddPartialPerson(t *testing.T) {
  db := MockDatabase{}
  librarian := Librarian{
    DatabasePtr: &db,
  }
  personPtr, err := librarian.add(Person{
    Nickname: "Johnny",
    FirstName: sql.NullString{String: "John", Valid: true,},
    Description: sql.NullString{String: "he seems like a nice guy", Valid: true,},
  })
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if personPtr.Nickname != "Johnny" {
    t.Errorf("expected Johnny got %v", personPtr.Nickname)
  }
  if personPtr.FirstName.String != "John" {
    t.Errorf("expected John got %v", personPtr.FirstName.String)
  }
  if personPtr.MiddleName.Valid != false {
    t.Errorf("expected null string got %v", personPtr.MiddleName.String)
  }
  if personPtr.LastName.Valid != false {
    t.Errorf("expected null string got %v", personPtr.LastName.String)
  }
  if personPtr.PhoneCountry.Valid != false {
    t.Errorf("expected null string got %v", personPtr.PhoneCountry.String)
  }
  if personPtr.PhoneArea.Valid != false {
    t.Errorf("expected null string got %v", personPtr.PhoneArea.String)
  }
  if personPtr.PhoneNumber.Valid != false {
    t.Errorf("expected null string got %v", personPtr.PhoneNumber.String)
  }
  if personPtr.Email.Valid != false {
    t.Errorf("expected null string got %v", personPtr.Email.String)
  }
  if personPtr.Birthdate.Valid != false {
    t.Errorf("expected null string got %v", personPtr.Birthdate.String)
  }
  if personPtr.Description.String != "he seems like a nice guy" {
    t.Errorf("expected \"he seems like a nice guy\" (without quotes) got %v", personPtr.Description.String)
  }
}

func TestUpdateId(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  personPtr, err := librarian.add(Person{
    Nickname: "God",
    LastName: sql.NullString{String: "Lawrence", Valid: true,},
  })
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  personPtr, err = librarian.update(Person{
    Id: personPtr.Id,
    Nickname: "Demon",
    FirstName: sql.NullString{String: "Satan", Valid: true,},
    LastName: sql.NullString{String: "", Valid: false,},
  })
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if personPtr.Nickname != "Demon" {
    t.Errorf("expected Demon got %v", personPtr.Nickname)
  }
  if personPtr.FirstName.String != "Satan" {
    t.Errorf("expected Satan got %v", personPtr.FirstName.String)
  }
  if personPtr.LastName.Valid != false {
    t.Errorf("expected null string got %v", personPtr.LastName.String)
  }
}

func TestUpdateIdNotExistId(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  personPtr, err := librarian.update(Person{
    Id: 1,
    Nickname: "Hello",
  })
  if err == nil {
    t.Errorf("expected error got nil")
  }
  if personPtr != nil {
    t.Errorf("expected nil got %v", personPtr)
  }
}

func TestSearchNickname(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  _, err := librarian.add(Person{
    Nickname: "Bullshit",
    FirstName: sql.NullString{String: "Lam", Valid: true,},
    MiddleName: sql.NullString{String: "Nha", Valid: true,},
    LastName: sql.NullString{String: "Tranh", Valid: true,},
    PhoneCountry: sql.NullString{String: "84", Valid: true,},
    PhoneArea: sql.NullString{String: "2", Valid: true,},
    PhoneNumber: sql.NullString{String: "111222333", Valid: true,},
    Email: sql.NullString{String: "nhatrang@gmail.com", Valid: true,},
    Birthdate: sql.NullString{String: "19970725", Valid: true,},
    Description: sql.NullString{String: "jerk", Valid: true,},
  })
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  _, err = librarian.add(Person{
    Nickname: "Bro BullshitAlpha",
    FirstName: sql.NullString{String: "What", Valid: true,},
  })
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  people, err := librarian.search("Bullshit")
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if len(people) != 2 {
    t.Errorf("expected 2 got %v", len(people))
  }
  if people[0].Nickname != "Bullshit" {
    t.Errorf("expected Bullshit got %v", people[0].Nickname)
  }
  if people[0].FirstName.String != "Lam" {
    t.Errorf("expected Lam got %v", people[0].FirstName.String)
  }
  if people[0].MiddleName.String != "Nha" {
    t.Errorf("expected Nha got %v", people[0].MiddleName.String)
  }
  if people[0].LastName.String != "Tranh" {
    t.Errorf("expected Tranh got %v", people[0].LastName.String)
  }
  if people[0].PhoneCountry.String != "84" {
    t.Errorf("expected 84 got %v", people[0].PhoneCountry.String)
  }
  if people[0].PhoneArea.String != "2" {
    t.Errorf("expected 2 got %v", people[0].PhoneArea.String)
  }
  if people[0].PhoneNumber.String != "111222333" {
    t.Errorf("expected 111222333 got %v", people[0].PhoneNumber.String)
  }
  if people[0].Email.String != "nhatrang@gmail.com" {
    t.Errorf("expected nhatrang@gmail.com got %v", people[0].Email.String)
  }
  if people[0].Birthdate.String != "19970725" {
    t.Errorf("expected 19970725 got %v", people[0].Birthdate.String)
  }
  if people[0].Description.String != "jerk" {
    t.Errorf("expected jerk got %v", people[0].Description.String)
  }
  if people[1].Nickname != "Bro BullshitAlpha" {
    t.Errorf("expected Bro BullshitAlpha got %v", people[1].Nickname)
  }
  if people[1].FirstName.String != "What" {
    t.Errorf("expected What got %v", people[1].FirstName.String)
  }
}

func TestSearchNicknameNotExist(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  people, err := librarian.search("Mom")
  if err != nil {
    t.Errorf("expected error to be nil got %v", err)
  }
  if len(people) != 0 {
    t.Errorf("expected 0 got %v", len(people))
  }
}
