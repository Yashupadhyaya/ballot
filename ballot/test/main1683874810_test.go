package main

import (
	"testing"
)

type Student struct {
	id   int
	name string
	age  int
}

func (s *Student) getId() int {
	return s.id
}

func (s *Student) getName() string {
	return s.name
}

func (s *Student) getAge() int {
	return s.age
}

var studentList = []Student{
	{1, "John", 20},
	{2, "Jane", 22},
	{3, "Doe", 19},
}

func fetchStudent(id int) (int, string, int) {
	for _, student := range studentList {
		if student.getId() == id {
			return student.getId(), student.getName(), student.getAge()
		}
	}
	return 0, "", 0
}

func TestFetchStudent(t *testing.T) {
	tests := []struct {
		id           int
		expectedId   int
		expectedName string
		expectedAge  int
	}{
		{1, 1, "John", 20},
		{2, 2, "Jane", 22},
		{3, 3, "Doe", 19},
		{4, 0, "", 0},
	}

	for _, test := range tests {
		id, name, age := fetchStudent(test.id)
		if id != test.expectedId || name != test.expectedName || age != test.expectedAge {
			t.Errorf("fetchStudent(%d) = (%d, %s, %d); expected (%d, %s, %d)", test.id, id, name, age, test.expectedId, test.expectedName, test.expectedAge)
		}
	}
}
