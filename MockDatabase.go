package main

import(
  "database/sql"
)

type MockDatabase struct {  
  Data []interface{}
}

func (dbPtr *MockDatabase) Query(query string, args ...interface{}) (IRows, error) {
  rows := MockRows{
    Data: dbPtr.Data,
    First: true,
  }
  return &rows, nil
}

func (dbPtr *MockDatabase) QueryRow(query string, args ...interface{}) (IRow) {
  row := MockRow{
    Data: dbPtr.Data,
  }
  return &row
}

func (dbPtr *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
  dbPtr.Data = args
  return nil, nil
}
