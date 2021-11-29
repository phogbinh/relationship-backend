package main

import(
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

func (librarian Librarian) search(nickname string) (Person, error) {
  return Person{}, nil
}
