package main

import(
  "database/sql"
)

type Person struct {
  Id int
  Nickname string
  FirstName sql.NullString
  MiddleName sql.NullString
  LastName sql.NullString
  PhoneCountry sql.NullString
  PhoneArea sql.NullString
  PhoneNumber sql.NullString
  Email sql.NullString
  Birthdate sql.NullString
  Description sql.NullString
}
