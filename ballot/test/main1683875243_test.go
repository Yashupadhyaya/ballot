package main

import (
	"testing"
)

type Person struct {
	age int
}

func (p *Person) getAge() int {
	return p.age
}

func TestGetAge(t *testing.T) {
	testCases := []struct {
		name     string
		person   *Person
		expected int
	}{
		{
			name:     "Positive age",
			person:   &Person{age: 25},
			expected: 25,
		},
		{
			name:     "Zero age",
			person:   &Person{age: 0},
			expected: 0,
		},
		{
			name:     "Negative age",
			person:   &Person{age: -5},
			expected: -5,
		},
		{
			name:     "Maximum age",
			person:   &Person{age: 2147483647},
			expected: 2147483647,
		},
		{
			name:     "Minimum age",
			person:   &Person{age: -2147483648},
			expected: -2147483648,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.person.getAge()
			if result != tc.expected {
				t.Errorf("Expected age %d, but got %d", tc.expected, result)
			}
		})
	}
}

