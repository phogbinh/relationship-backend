package main

import(
  "database/sql"
)

type MockDatabase struct {  
  Data []interface{}
}

func (databasePtr *MockDatabase) Query(query string, args ...interface{}) (IRows, error) {
  rows := MockRows{
    Data: databasePtr.Data,
    First: true,
  }
  return &rows, nil
}

func (databasePtr *MockDatabase) QueryRow(query string, args ...interface{}) (IRow) {
  row := MockRow{
    Data: databasePtr.Data,
  }
  return &row
}

func (databasePtr *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
  databasePtr.Data = args
  return nil, nil
}
