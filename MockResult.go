package main

type MockResult struct {
}

func (result MockResult) LastInsertId() (int64, error) {
  return 0, nil
}

func (result MockResult) RowsAffected() (int64, error) {
  return 0, nil
}
