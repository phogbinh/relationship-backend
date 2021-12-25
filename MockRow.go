package main

import(
  "database/sql"
  "reflect"
)

type MockRow struct {
  Data []interface{}
}

func (rowPtr *MockRow) Scan(dest ...interface{}) (error) {
  if len(rowPtr.Data) == 0 {
    return sql.ErrNoRows
  }
  for i := 1; i < len(dest); i++ {
    reflect.ValueOf(dest[i]).Elem().Set(reflect.ValueOf(rowPtr.Data[i - 1]))
  }
  return nil
}
