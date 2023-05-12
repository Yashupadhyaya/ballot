package main

import (
	"testing"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

var studentList []Student

func addStudent(id int, name string, age int) {
	studentList = append(studentList, Student{ID: id, Name: name, Age: age})
}

func TestAddStudent(t *testing.T) {
	t.Run("Adding valid student", func(t *testing.T) {
		addStudent(1, "John Doe", 20)
		if len(studentList) != 1 {
			t.Errorf("Expected student list length to be 1, got %d", len(studentList))
		}
	})

	t.Run("Adding student with negative age", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic for negative age")
			}
		}()
		addStudent(2, "Jane Doe", -5)
	})

	t.Run("Adding student with empty name", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic for empty name")
			}
		}()
		addStudent(3, "", 25)
	})

	t.Run("Adding student with duplicate ID", func(t *testing.T) {
		addStudent(4, "Alice", 30)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic for duplicate ID")
			}
		}()
		addStudent(4, "Bob", 35)
	})
}
