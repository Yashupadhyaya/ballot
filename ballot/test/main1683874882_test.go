package main

import (
	"testing"
)

type Student struct {
	id int
}

func (s Student) GetId() int {
	return s.id
}

var studentList = []Student{
	{1},
	{2},
	{3},
	{4},
	{5},
}

func isPresent(id int) (bool, error) {
	for _, student := range studentList {
		if student.GetId() == id {
			return true, nil
		}
	}
	return false, nil
}

func TestIsPresent(t *testing.T) {
	tests := []struct {
		id          int
		expected    bool
		description string
	}{
		{1, true, "Student with id 1 is present"},
		{2, true, "Student with id 2 is present"},
		{3, true, "Student with id 3 is present"},
		{4, true, "Student with id 4 is present"},
		{5, true, "Student with id 5 is present"},
		{6, false, "Student with id 6 is not present"},
		{7, false, "Student with id 7 is not present"},
		{-1, false, "Student with negative id is not present"},
	}

	for _, test := range tests {
		result, err := isPresent(test.id)
		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}
		if result != test.expected {
			t.Errorf("Test failed for %s: expected %v, got %v", test.description, test.expected, result)
		}
	}
}
