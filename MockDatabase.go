package main

import(
  "database/sql"
  "errors"
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
  if databasePtr.Data == nil {
    return new(MockRow)
  }
  row := MockRow {
    Data: databasePtr.Data[0],
  }
  return &row
}

func (databasePtr *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
  if query[:6] == "UPDATE" {
    if databasePtr.Data == nil {
      return new(MockResult), nil
    }
    databasePtr.Data = nil // workaround update by removing all elements to add new
  } else if query[:6] == "DELETE" {
    if databasePtr.Data == nil {
      return nil, errors.New("")
    }
  }
  databasePtr.Data = append(databasePtr.Data, args)
  return new(MockResult), nil
}
