package main

import(
  "database/sql"
)

type IDatabase interface {
  Query(query string, args ...interface{}) (IRows, error)
  QueryRow(query string, args ...interface{}) (IRow)
  Exec(query string, args ...interface{}) (sql.Result, error)
}
