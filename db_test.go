package main

import "testing"

func TestGetRecordsFromDB(t *testing.T) {
	rec, err := getRecordsFromDB()
	if err != nil {
		t.Error(err)
	}

	if len(rec) < 1 {
		t.Errorf("Expected >= 1 records got %d", len(rec))
	}
}
