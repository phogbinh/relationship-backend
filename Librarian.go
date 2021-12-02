package main

import(
  "database/sql"
  "fmt"
)

type Librarian struct {
  DatabasePtr *MockDatabase
}

func (librarian *Librarian) add(person Person) (error) {
  _, err := librarian.DatabasePtr.Exec("INSERT INTO person(nickname, first_name, middle_name, last_name, phone_country, phone_area, phone_number, email, birthdate, description) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", person.Nickname, person.FirstName, person.MiddleName, person.LastName, person.PhoneCountry, person.PhoneArea, person.PhoneNumber, person.Email, person.Birthdate, person.Description)
  if err != nil {
    return fmt.Errorf("add: %v", err)
  }
  return nil
}

func (librarian Librarian) search(nickname string) (*Person, error) {
  var id int
  personPtr := new(Person)
  err := librarian.DatabasePtr.QueryRow("SELECT * FROM person WHERE nickname = ?", nickname).Scan(&id, &personPtr.Nickname, &personPtr.FirstName, &personPtr.MiddleName, &personPtr.LastName, &personPtr.PhoneCountry, &personPtr.PhoneArea, &personPtr.PhoneNumber, &personPtr.Email, &personPtr.Birthdate, &personPtr.Description)
  if err != nil {
    if err == sql.ErrNoRows {
      return nil, fmt.Errorf("search %v: unknown nickname", nickname)
    }
    return nil, fmt.Errorf("search %v: %v", nickname, err)
  }
  return personPtr, nil
}
