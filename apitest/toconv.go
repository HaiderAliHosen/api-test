package apitest

import "testing"

// TestAdd -- to addd post
func TestAdd(t *testing.T) {
	serv := New()
	serv.Add(Item{})
	if len(serv.Items) == 0 {
		t.Errorf("Item was not added")
	}
}

// TestGetAll for getall
func TestGetAll(t *testing.T) {
	serv := New()
	serv.Add(Item{})
	results := serv.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
