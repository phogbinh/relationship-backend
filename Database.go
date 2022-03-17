package main

import(
  "database/sql"
)

type Database struct {
  SqlDatabasePtr *sql.DB
}

func (databasePtr *Database) Query(query string, args ...interface{}) (IRows, error) {
  return databasePtr.SqlDatabasePtr.Query(query, args...)
}

func (databasePtr *Database) QueryRow(query string, args ...interface{}) (IRow) {
  return databasePtr.SqlDatabasePtr.QueryRow(query, args...)
}

func (databasePtr *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
  return databasePtr.SqlDatabasePtr.Exec(query, args...)
}
