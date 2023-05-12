package main

import (
	"testing"
)

type Student struct {
	id   int
	name string
	age  int
}

func NewStudent(id int, name string, age int) *Student {
	return &Student{id: id, name: name, age: age}
}

func TestNewStudent(t *testing.T) {
	tests := []struct {
		id        int
		name      string
		age       int
		expectErr bool
	}{
		{1, "John Doe", 25, false},
		{2, "Jane Doe", 30, false},
		{-1, "Invalid ID", 20, true},
		{3, "", 22, true},
		{4, "Negative Age", -5, true},
	}

	for _, test := range tests {
		student, err := NewStudent(test.id, test.name, test.age)
		if test.expectErr {
			if err == nil {
				t.Errorf("Expected error for input: %v, but got no error", test)
			}
		} else {
			if err != nil {
				t.Errorf("Expected no error for input: %v, but got error: %v", test, err)
			}
			if student.id != test.id || student.name != test.name || student.age != test.age {
				t.Errorf("Expected student with id: %d, name: %s, age: %d, but got id: %d, name: %s, age: %d",
					test.id, test.name, test.age, student.id, student.name, student.age)
			}
		}
	}
}
