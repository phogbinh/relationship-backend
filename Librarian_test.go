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
  err := librarian.add(Person{
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
  if person.FirstName.String != "Lam" {
    t.Errorf("expected Lam got %v", person.FirstName.String)
  }
  if person.MiddleName.String != "Nha" {
    t.Errorf("expected Nha got %v", person.MiddleName.String)
  }
  if person.LastName.String != "Tranh" {
    t.Errorf("expected Tranh got %v", person.LastName.String)
  }
  if person.PhoneCountry.String != "84" {
    t.Errorf("expected 84 got %v", person.PhoneCountry.String)
  }
  if person.PhoneArea.String != "2" {
    t.Errorf("expected 2 got %v", person.PhoneArea.String)
  }
  if person.PhoneNumber.String != "111222333" {
    t.Errorf("expected 111222333 got %v", person.PhoneNumber.String)
  }
  if person.Email.String != "nhatrang@gmail.com" {
    t.Errorf("expected nhatrang@gmail.com got %v", person.Email.String)
  }
  if person.Birthdate.String != "19970725" {
    t.Errorf("expected 19970725 got %v", person.Birthdate.String)
  }
  if person.Description.String != "jerk" {
    t.Errorf("expected jerk got %v", person.Description.String)
  }
}

func TestAddPartialPerson(t *testing.T) {
  db := MockDatabase{}
  librarian := Librarian{
    DatabasePtr: &db,
  }
  err := librarian.add(Person{
    Nickname: "Johnny",
    FirstName: sql.NullString{String: "John", Valid: true,},
    Description: sql.NullString{String: "he seems like a nice guy", Valid: true,},
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
  if person.FirstName.String != "John" {
    t.Errorf("expected John got %v", person.FirstName.String)
  }
  if person.MiddleName.String != "" {
    t.Errorf("expected empty string got %v", person.MiddleName.String)
  }
  if person.LastName.String != "" {
    t.Errorf("expected empty string got %v", person.LastName.String)
  }
  if person.PhoneCountry.String != "" {
    t.Errorf("expected empty string got %v", person.PhoneCountry.String)
  }
  if person.PhoneArea.String != "" {
    t.Errorf("expected empty string got %v", person.PhoneArea.String)
  }
  if person.PhoneNumber.String != "" {
    t.Errorf("expected empty string got %v", person.PhoneNumber.String)
  }
  if person.Email.String != "" {
    t.Errorf("expected empty string got %v", person.Email.String)
  }
  if person.Birthdate.String != "" {
    t.Errorf("expected empty string got %v", person.Birthdate.String)
  }
  if person.Description.String != "he seems like a nice guy" {
    t.Errorf("expected \"he seems like a nice guy\" (without quotes) got %v", person.Description.String)
  }
}

func TestSearchNickname(t *testing.T) {
  librarian := Librarian{
    DatabasePtr: new(MockDatabase),
  }
  err := librarian.add(Person{
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
  err = librarian.add(Person{
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

func TestSearchNotExist(t *testing.T) {
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
