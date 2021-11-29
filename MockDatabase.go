package main

import(
  "database/sql"
)

type MockDatabase struct {  
  Data []interface{}
}

func (dbPtr *MockDatabase) Query(query string, args ...interface{}) (IRows, error) {
  rowsPtr := new(MockRows)
  rowsPtr.Data = dbPtr.Data
  rowsPtr.First = true
  return rowsPtr, nil
}

func (dbPtr *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
  dbPtr.Data = args
  return nil, nil
}
