package main

import(
  "database/sql"
  "fmt"
)

type Librarian struct {
  DatabasePtr IDatabase
}

func (librarian *Librarian) add(person Person) (*Person, error) {
  result, err := librarian.DatabasePtr.Exec("INSERT INTO person(nickname, first_name, middle_name, last_name, phone_country, phone_area, phone_number, email, birthdate, description) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", person.Nickname, person.FirstName, person.MiddleName, person.LastName, person.PhoneCountry, person.PhoneArea, person.PhoneNumber, person.Email, person.Birthdate, person.Description)
  if err != nil {
    return nil, fmt.Errorf("add: %v", err)
  }
  id, err := result.LastInsertId()
  if err != nil {
    return nil, fmt.Errorf("get id after add: %v", err)
  }
  err = librarian.DatabasePtr.QueryRow("SELECT * FROM person WHERE id = ?", id).Scan(&person.Id, &person.Nickname, &person.FirstName, &person.MiddleName, &person.LastName, &person.PhoneCountry, &person.PhoneArea, &person.PhoneNumber, &person.Email, &person.Birthdate, &person.Description)
  if err != nil {
    if err == sql.ErrNoRows {
      return nil, fmt.Errorf("get person %d: unknown person", id)
    }
    return nil, fmt.Errorf("get person %d: %v", id, err)
  }
  return &person, nil
}

func (librarian Librarian) update(person Person) (*Person, error) {
  _, err := librarian.DatabasePtr.Exec("UPDATE person SET nickname = ?, first_name = ?, middle_name = ?, last_name = ?, phone_country = ?, phone_area = ?, phone_number = ?, email = ?, birthdate = ?, description = ? WHERE id = ?", person.Nickname, person.FirstName, person.MiddleName, person.LastName, person.PhoneCountry, person.PhoneArea, person.PhoneNumber, person.Email, person.Birthdate, person.Description, person.Id)
  if err != nil {
    return nil, fmt.Errorf("update: %v", err)
  }
  id := person.Id
  err = librarian.DatabasePtr.QueryRow("SELECT * FROM person WHERE id = ?", id).Scan(&person.Id, &person.Nickname, &person.FirstName, &person.MiddleName, &person.LastName, &person.PhoneCountry, &person.PhoneArea, &person.PhoneNumber, &person.Email, &person.Birthdate, &person.Description)
  if err != nil {
    if err == sql.ErrNoRows {
      return nil, fmt.Errorf("get person %d: unknown person", id)
    }
    return nil, fmt.Errorf("get person %d: %v", id, err)
  }
  return &person, nil
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
