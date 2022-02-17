package main

import(
  "fmt"
)

type Librarian struct {
  DatabasePtr IDatabase
}

func (librarian *Librarian) add(person Person) (error) {
  _, err := librarian.DatabasePtr.Exec("INSERT INTO person(nickname, first_name, middle_name, last_name, phone_country, phone_area, phone_number, email, birthdate, description) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", person.Nickname, person.FirstName, person.MiddleName, person.LastName, person.PhoneCountry, person.PhoneArea, person.PhoneNumber, person.Email, person.Birthdate, person.Description)
  if err != nil {
    return fmt.Errorf("add: %v", err)
  }
  return nil
}

func (librarian Librarian) search(nickname string) ([]Person, error) {
  rowsPtr, err := librarian.DatabasePtr.Query("SELECT * FROM person WHERE nickname REGEXP ?", nickname)
  if err != nil {
    return nil, err
  }
  defer rowsPtr.Close()
  people := []Person{} // prevent null on json marshal
  for rowsPtr.Next() {
    var person Person
    err = rowsPtr.Scan(&person.Id, &person.Nickname, &person.FirstName, &person.MiddleName, &person.LastName, &person.PhoneCountry, &person.PhoneArea, &person.PhoneNumber, &person.Email, &person.Birthdate, &person.Description)
    if err != nil {
      return people, err
    }
    people = append(people, person)
  }
  err = rowsPtr.Err()
  if err != nil {
    return people, err
  }
  return people, nil
}
