package main

import(
  "database/sql"
  "reflect"
)

type MockRows struct {
  Data []interface{}
  First bool
}

func (rowsPtr *MockRows) Close() (error) {
  return nil
}

func (rowsPtr *MockRows) Err() (error) {
  return nil
}

func (rowsPtr *MockRows) Next() (bool) {
  if rowsPtr.First {
    rowsPtr.First = false
    return true
  }
  return false
}

func (rowsPtr *MockRows) Scan(dest ...interface{}) (error) {
  if len(rowsPtr.Data) == 0 {
    return sql.ErrNoRows
  }
  for i := 1; i < len(dest); i++ {
    reflect.ValueOf(dest[i]).Elem().Set(reflect.ValueOf(rowsPtr.Data[i - 1]))
  }
  return nil
}
