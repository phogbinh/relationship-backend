package main

type IRows interface {
  Close() error
  Err() error
  Next() bool
  Scan(dest ...interface{}) error
}
