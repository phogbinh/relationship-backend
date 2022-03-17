package main

import(
  "database/sql"
)

type MockDatabase struct {  
  Data [][]interface{}
}

func (databasePtr *MockDatabase) Query(query string, args ...interface{}) (IRows, error) {
  rows := MockRows{
    Data: databasePtr.Data,
    CurrentIndex: 0,
  }
  return &rows, nil
}

func (databasePtr *MockDatabase) QueryRow(query string, args ...interface{}) (IRow) {
  row := MockRow {
    Data: databasePtr.Data[0],
  }
  return &row
}

func (databasePtr *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
  databasePtr.Data = append(databasePtr.Data, args)
  return new(MockResult), nil
}
