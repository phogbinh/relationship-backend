package main

type IRow interface {
  Scan(dest ...interface{}) (error)
}
