package main

import(
  "reflect"
)

type MockRow struct {
  Data []interface{}
}

func (rowPtr *MockRow) Scan(dest ...interface{}) (error) {
  for i := 1; i < len(dest); i++ {
    reflect.ValueOf(dest[i]).Elem().Set(reflect.ValueOf(rowPtr.Data[i - 1]))
  }
  return nil
}
