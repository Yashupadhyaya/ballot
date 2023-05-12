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

func (s *Student) setName(name string) {
	s.name = name
}

func (s *Student) setAge(age int) {
	s.age = age
}

var studentList = []Student{
	{1, "John", 20},
	{2, "Jane", 22},
	{3, "Mark", 24},
}

func updateStudent(id int, name string, age int) bool {
	for i, student := range studentList {
		if student.getId() == id {
			studentList[i].setName(name)
			studentList[i].setAge(age)
			return true
		}
	}
	return false
}

func TestUpdateStudent(t *testing.T) {
	// Test case 1: Update an existing student
	updated := updateStudent(1, "John Doe", 21)
	if !updated {
		t.Error("Expected to update an existing student")
	}
	if studentList[0].getName() != "John Doe" || studentList[0].getAge() != 21 {
		t.Error("Expected updated student to have new name and age")
	}

	// Test case 2: Try to update a non-existent student
	updated = updateStudent(4, "Non Existent", 25)
	if updated {
		t.Error("Expected not to update a non-existent student")
	}

	// Test case 3: Update student with edge case values
	updated = updateStudent(2, "", 0)
	if !updated {
		t.Error("Expected to update an existing student with edge case values")
	}
	if studentList[1].getName() != "" || studentList[1].getAge() != 0 {
		t.Error("Expected updated student to have empty name and age 0")
	}
}
