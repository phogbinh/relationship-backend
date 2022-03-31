package main

import(
  "database/sql"
)

type Person struct {
  Id int64
  Nickname string `json:"nickname"`
  FirstName sql.NullString `json:"firstName"`
  MiddleName sql.NullString `json:"middleName"`
  LastName sql.NullString `json:"lastName"`
  PhoneCountry sql.NullString `json:"phoneCountry"`
  PhoneArea sql.NullString `json:"phoneArea"`
  PhoneNumber sql.NullString `json:"phoneNumber"`
  Email sql.NullString `json:"email"`
  Birthdate sql.NullString `json:"birthdate"`
  Description sql.NullString `json:"description"`
}
