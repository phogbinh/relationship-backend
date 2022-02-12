package main

import(
  "database/sql"
  "reflect"
)

type MockRows struct {
  Data [][]interface{}
  CurrentIndex int
}

func (rowsPtr *MockRows) Close() (error) {
  return nil
}

func (rowsPtr *MockRows) Err() (error) {
  return nil
}

func (rowsPtr *MockRows) Next() (bool) {
  return rowsPtr.CurrentIndex < len(rowsPtr.Data)
}

func (rowsPtr *MockRows) Scan(dest ...interface{}) (error) {
  if len(rowsPtr.Data[rowsPtr.CurrentIndex]) == 0 {
    return sql.ErrNoRows
  }
  for i := 1; i < len(dest); i++ {
    reflect.ValueOf(dest[i]).Elem().Set(reflect.ValueOf(rowsPtr.Data[rowsPtr.CurrentIndex][i - 1]))
  }
  rowsPtr.CurrentIndex++
  return nil
}