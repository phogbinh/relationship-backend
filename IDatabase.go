package main

import(
  "database/sql"
)

type IDatabase interface {
  Query(query string, args ...interface{}) (IRows, error)
  Exec(query string, args ...interface{}) (sql.Result, error)
}
