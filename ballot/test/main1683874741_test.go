package main

import (
	"testing"
)

func TestDeleteStudent(t *testing.T) {
	// Setup
	studentList := []Student{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	tests := []struct {
		name          string
		idToDelete    int
		expectedList  []Student
		expectedError error
	}{
		{
			name:       "Delete existing student",
			idToDelete: 2,
			expectedList: []Student{
				{ID: 1, Name: "Alice"},
				{ID: 3, Name: "Charlie"},
			},
			expectedError: nil,
		},
		{
			name:       "Delete non-existing student",
			idToDelete: 4,
			expectedList: []Student{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 3, Name: "Charlie"},
			},
			expectedError: fmt.Errorf("Student not found"),
		},
		{
			name:       "Delete student with negative ID",
			idToDelete: -1,
			expectedList: []Student{
				{ID: 1, Name: "Alice"},
				{ID: 2, Name: "Bob"},
				{ID: 3, Name: "Charlie"},
			},
			expectedError: fmt.Errorf("Invalid ID"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Execute
			err := deleteStudent(test.idToDelete)

			// Verify
			if !reflect.DeepEqual(studentList, test.expectedList) {
				t.Errorf("Expected student list: %v, got: %v", test.expectedList, studentList)
			}

			if !errors.Is(err, test.expectedError) {
				t.Errorf("Expected error: %v, got: %v", test.expectedError, err)
			}
		})
	}
}

