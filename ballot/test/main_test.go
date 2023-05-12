package main

import (
	"testing"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

var studentList = []Student{
	{1, "John Doe", 20, "A"},
	{2, "Jane Doe", 22, "B"},
	{3, "Alice", 19, "C"},
	{4, "Bob", 21, "D"},
}

func TestAddStudent(t *testing.T) {
	newStudent := Student{5, "Charlie", 23, "E"}
	studentList = append(studentList, newStudent)

	if len(studentList) != 5 {
		t.Errorf("Expected student list length to be 5, got %d", len(studentList))
	}

	if studentList[4] != newStudent {
		t.Errorf("Expected last student to be %v, got %v", newStudent, studentList[4])
	}
}

func TestUpdateStudent(t *testing.T) {
	updatedStudent := Student{1, "John Doe", 20, "A+"}
	for i, student := range studentList {
		if student.ID == updatedStudent.ID {
			studentList[i] = updatedStudent
			break
		}
	}

	if studentList[0] != updatedStudent {
		t.Errorf("Expected updated student to be %v, got %v", updatedStudent, studentList[0])
	}
}

func TestDeleteStudent(t *testing.T) {
	deletedID := 3
	for i, student := range studentList {
		if student.ID == deletedID {
			studentList = append(studentList[:i], studentList[i+1:]...)
			break
		}
	}

	for _, student := range studentList {
		if student.ID == deletedID {
			t.Errorf("Expected student with ID %d to be deleted", deletedID)
		}
	}
}

func TestFetchStudent(t *testing.T) {
	fetchedID := 2
	var fetchedStudent Student
	for _, student := range studentList {
		if student.ID == fetchedID {
			fetchedStudent = student
			break
		}
	}

	if fetchedStudent.ID != fetchedID {
		t.Errorf("Expected fetched student ID to be %d, got %d", fetchedID, fetchedStudent.ID)
	}
}

func TestIsPresent(t *testing.T) {
	searchID := 4
	present := false
	for _, student := range studentList {
		if student.ID == searchID {
			present = true
			break
		}
	}

	if !present {
		t.Errorf("Expected student with ID %d to be present", searchID)
	}
}

func TestListStudent(t *testing.T) {
	if len(studentList) == 0 {
		t.Errorf("Expected student list to be non-empty")
	}
}

func TestInvalidInput(t *testing.T) {
	// Test for invalid input, such as negative student ID, empty name, etc.
	// This test case depends on the implementation of error handling in the main function.
}
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
	studentList = append(studentList, Student{id, name, age})
}

func TestAddStudent(t *testing.T) {
	// Test case 1: Adding a valid student
	addStudent(1, "John Doe", 25)
	if len(studentList) != 1 {
		t.Errorf("Expected studentList length to be 1, got %d", len(studentList))
	}
	if studentList[0].ID != 1 || studentList[0].Name != "John Doe" || studentList[0].Age != 25 {
		t.Errorf("Expected student with ID 1, Name 'John Doe', Age 25, got %v", studentList[0])
	}

	// Test case 2: Adding another valid student
	addStudent(2, "Jane Doe", 22)
	if len(studentList) != 2 {
		t.Errorf("Expected studentList length to be 2, got %d", len(studentList))
	}
	if studentList[1].ID != 2 || studentList[1].Name != "Jane Doe" || studentList[1].Age != 22 {
		t.Errorf("Expected student with ID 2, Name 'Jane Doe', Age 22, got %v", studentList[1])
	}

	// Test case 3: Adding a student with a negative age
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic when adding a student with a negative age")
		}
	}()
	addStudent(3, "Invalid Student", -5)
}

func TestMain(m *testing.M) {
	// Setup
	studentList = make([]Student, 0)

	// Run tests
	m.Run()

	// Teardown
	studentList = nil
}
package main

import (
	"testing"
)

type Student struct {
	id   int
	age  int
	name string
}

var studentList = []Student{
	{1, 20, "Alice"},
	{2, 22, "Bob"},
	{3, 21, "Charlie"},
}

func listStudent() {
	for _, student := range studentList {
		println(student.id, student.age, student.name)
	}
}

func TestListStudent(t *testing.T) {
	// Save original studentList
	originalStudentList := studentList

	// Test case 1: Normal student list
	studentList = []Student{
		{1, 20, "Alice"},
		{2, 22, "Bob"},
		{3, 21, "Charlie"},
	}
	listStudent()

	// Test case 2: Empty student list
	studentList = []Student{}
	listStudent()

	// Test case 3: Student list with invalid data
	studentList = []Student{
		{1, -1, "Alice"},
		{2, 22, ""},
	}
	listStudent()

	// Test case 4: Student list with duplicate data
	studentList = []Student{
		{1, 20, "Alice"},
		{1, 20, "Alice"},
	}
	listStudent()

	// Restore original studentList
	studentList = originalStudentList
}
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

func (s *Student) setName(name string) {
	s.name = name
}

func (s *Student) setAge(age int) {
	s.age = age
}

var studentList = []Student{
	{1, "John", 20},
	{2, "Jane", 22},
	{3, "Bob", 19},
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
	tests := []struct {
		id       int
		name     string
		age      int
		expected bool
	}{
		{1, "John Doe", 21, true},
		{2, "Jane Doe", 23, true},
		{3, "Bob Smith", 20, true},
		{4, "Alice", 25, false},
		{-1, "Invalid", 30, false},
	}

	for _, test := range tests {
		result := updateStudent(test.id, test.name, test.age)
		if result != test.expected {
			t.Errorf("updateStudent(%d, %s, %d) = %v; expected %v", test.id, test.name, test.age, result, test.expected)
		}
	}
}
package main

import (
	"testing"
)

type Student struct {
	ID int
}

type StudentList []Student

func (s *StudentList) deleteStudent(id int) bool {
	for i, student := range *s {
		if student.ID == id {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return true
		}
	}
	return false
}

func TestDeleteStudent(t *testing.T) {
	studentList := StudentList{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
	}

	tests := []struct {
		name     string
		id       int
		expected bool
	}{
		{"Valid ID", 2, true},
		{"Invalid ID", 5, false},
		{"Negative ID", -1, false},
		{"Zero ID", 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := studentList.deleteStudent(test.id)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
package main

import (
	"testing"
)

type Student struct {
	id   int
	name string
	age  int
}

func (s Student) getId() int {
	return s.id
}

func (s Student) getName() string {
	return s.name
}

func (s Student) getAge() int {
	return s.age
}

var studentList = []Student{
	{1, "John", 20},
	{2, "Jane", 22},
	{3, "Doe", 24},
}

func fetchStudent(id int) (int, string, int, bool) {
	for _, student := range studentList {
		if student.getId() == id {
			return student.getId(), student.getName(), student.getAge(), true
		}
	}
	return 0, "", 0, false
}

func TestFetchStudent(t *testing.T) {
	tests := []struct {
		id           int
		expectedId   int
		expectedName string
		expectedAge  int
		expectedOk   bool
	}{
		{1, 1, "John", 20, true},
		{2, 2, "Jane", 22, true},
		{3, 3, "Doe", 24, true},
		{4, 0, "", 0, false},
		{-1, 0, "", 0, false},
	}

	for _, test := range tests {
		id, name, age, ok := fetchStudent(test.id)
		if id != test.expectedId || name != test.expectedName || age != test.expectedAge || ok != test.expectedOk {
			t.Errorf("fetchStudent(%d) = (%d, %s, %d, %t); want (%d, %s, %d, %t)", test.id, id, name, age, ok, test.expectedId, test.expectedName, test.expectedAge, test.expectedOk)
		}
	}
}
package main

import (
	"testing"
)

type Student struct {
	id int
}

func (s Student) getId() int {
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
		if student.getId() == id {
			return true, nil
		}
	}
	return false, nil
}

func TestIsPresent(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, false},
		{7, false},
		{-1, false},
	}

	for _, test := range tests {
		result, err := isPresent(test.id)
		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}
		if result != test.expected {
			t.Errorf("isPresent(%d) = %v; expected %v", test.id, result, test.expected)
		}
	}
}
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
	return &Student{
		id:   id,
		name: name,
		age:  age,
	}
}

func TestNewStudent(t *testing.T) {
	tests := []struct {
		id       int
		name     string
		age      int
		expected *Student
	}{
		{1, "John Doe", 20, &Student{1, "John Doe", 20}},
		{2, "Jane Doe", 25, &Student{2, "Jane Doe", 25}},
		{-1, "Invalid ID", 30, &Student{-1, "Invalid ID", 30}},
		{3, "", 35, &Student{3, "", 35}},
		{4, "Max Age", 200, &Student{4, "Max Age", 200}},
	}

	for _, test := range tests {
		result := NewStudent(test.id, test.name, test.age)
		if result.id != test.expected.id || result.name != test.expected.name || result.age != test.expected.age {
			t.Errorf("Expected student with id %d, name %s, and age %d, but got id %d, name %s, and age %d", test.expected.id, test.expected.name, test.expected.age, result.id, result.name, result.age)
		}
	}
}

func TestNewStudentErrorHandling(t *testing.T) {
	tests := []struct {
		id       int
		name     string
		age      int
		expected *Student
	}{
		{0, "Zero ID", 20, nil},
		{5, "Negative Age", -5, nil},
		{6, "Too Long Name", 30, nil},
	}

	for _, test := range tests {
		result := NewStudent(test.id, test.name, test.age)
		if result != test.expected {
			t.Errorf("Expected error for student with id %d, name %s, and age %d, but got a valid student", test.id, test.name, test.age)
		}
	}
}
package main

import (
	"testing"
)

type Product struct {
	id       int
	name     string
	price    float64
	rating   float64
	category string
}

func (p *Product) getId() int {
	return p.id
}

func TestGetId(t *testing.T) {
	testCases := []struct {
		name     string
		product  Product
		expected int
	}{
		{
			name:     "Valid ID",
			product:  Product{id: 1, name: "Test Product", price: 10.0, rating: 4.5, category: "Electronics"},
			expected: 1,
		},
		{
			name:     "Zero ID",
			product:  Product{id: 0, name: "Test Product", price: 10.0, rating: 4.5, category: "Electronics"},
			expected: 0,
		},
		{
			name:     "Negative ID",
			product:  Product{id: -1, name: "Test Product", price: 10.0, rating: 4.5, category: "Electronics"},
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.product.getId()
			if result != tc.expected {
				t.Errorf("Expected ID %d, but got %d", tc.expected, result)
			}
		})
	}
}
package main

import (
	"testing"
)

type Product struct {
	id       int
	keywords []string
	price    float64
	rating   float64
	category string
}

type ShoppingCart struct {
	items []Product
}

func (p *Product) setId(id int) {
	p.id = id
}

func (s *ShoppingCart) addItem(product Product) {
	s.items = append(s.items, product)
}

func (s *ShoppingCart) totalCost() float64 {
	var total float64
	for _, item := range s.items {
		total += item.price
	}
	return total
}

func TestSetId(t *testing.T) {
	product := Product{}
	product.setId(1)

	if product.id != 1 {
		t.Errorf("Expected id to be 1, but got %d", product.id)
	}
}

func TestAddItem(t *testing.T) {
	shoppingCart := ShoppingCart{}
	product := Product{id: 1, price: 10.0}

	shoppingCart.addItem(product)

	if len(shoppingCart.items) != 1 {
		t.Errorf("Expected 1 item in the shopping cart, but got %d", len(shoppingCart.items))
	}
}

func TestTotalCost(t *testing.T) {
	shoppingCart := ShoppingCart{}
	product1 := Product{id: 1, price: 10.0}
	product2 := Product{id: 2, price: 20.0}

	shoppingCart.addItem(product1)
	shoppingCart.addItem(product2)

	total := shoppingCart.totalCost()

	if total != 30.0 {
		t.Errorf("Expected total cost to be 30.0, but got %f", total)
	}
}

func TestTotalCostEmptyCart(t *testing.T) {
	shoppingCart := ShoppingCart{}

	total := shoppingCart.totalCost()

	if total != 0.0 {
		t.Errorf("Expected total cost to be 0.0, but got %f", total)
	}
}
package main

import (
	"testing"
)

type Product struct {
	name  string
	price float64
	rating int
	category string
}

func (p *Product) getName() string {
	return p.name
}

func TestGetName(t *testing.T) {
	testCases := []struct {
		name     string
		product  Product
		expected string
	}{
		{
			name:     "Success: Get name of a product",
			product:  Product{name: "Laptop", price: 1000, rating: 5, category: "Electronics"},
			expected: "Laptop",
		},
		{
			name:     "Success: Get name of a product with special characters",
			product:  Product{name: "Laptop@#$%^&*()", price: 1000, rating: 5, category: "Electronics"},
			expected: "Laptop@#$%^&*()",
		},
		{
			name:     "Success: Get name of a product with empty name",
			product:  Product{name: "", price: 1000, rating: 5, category: "Electronics"},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.product.getName()
			if result != tc.expected {
				t.Errorf("Expected: %s, got: %s", tc.expected, result)
			}
		})
	}
}
package main

import (
	"testing"
)

type Product struct {
	name string
}

func (p *Product) setName(name string) {
	p.name = name
}

func TestSetName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid name",
			input:    "Laptop",
			expected: "Laptop",
		},
		{
			name:     "Empty name",
			input:    "",
			expected: "",
		},
		{
			name:     "Name with special characters",
			input:    "Laptop@123",
			expected: "Laptop@123",
		},
		{
			name:     "Name with spaces",
			input:    "Laptop Pro",
			expected: "Laptop Pro",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := &Product{}
			p.setName(test.input)

			if p.name != test.expected {
				t.Errorf("Expected name to be '%s', got '%s'", test.expected, p.name)
			}
		})
	}
}
package main

import (
	"testing"
)

type User struct {
	age int
}

func (u *User) getAge() int {
	return u.age
}

func TestGetAge(t *testing.T) {
	t.Run("Success: Valid age", func(t *testing.T) {
		user := &User{age: 25}
		age := user.getAge()

		if age != 25 {
			t.Errorf("Expected age to be 25, but got %d", age)
		}
	})

	t.Run("Success: Zero age", func(t *testing.T) {
		user := &User{age: 0}
		age := user.getAge()

		if age != 0 {
			t.Errorf("Expected age to be 0, but got %d", age)
		}
	})

	t.Run("Failure: Negative age", func(t *testing.T) {
		user := &User{age: -5}
		age := user.getAge()

		if age >= 0 {
			t.Errorf("Expected age to be negative, but got %d", age)
		}
	})
}
package main

import (
	"testing"
)

type Person struct {
	age int
}

func (p *Person) setAge(age int) {
	p.age = age
}

func TestSetAge(t *testing.T) {
	t.Run("Positive age", func(t *testing.T) {
		p := &Person{}
		p.setAge(25)
		if p.age != 25 {
			t.Errorf("Expected age to be 25, got %d", p.age)
		}
	})

	t.Run("Zero age", func(t *testing.T) {
		p := &Person{}
		p.setAge(0)
		if p.age != 0 {
			t.Errorf("Expected age to be 0, got %d", p.age)
		}
	})

	t.Run("Negative age", func(t *testing.T) {
		p := &Person{}
		p.setAge(-5)
		if p.age != -5 {
			t.Errorf("Expected age to be -5, got %d", p.age)
		}
	})

	t.Run("Large age", func(t *testing.T) {
		p := &Person{}
		p.setAge(150)
		if p.age != 150 {
			t.Errorf("Expected age to be 150, got %d", p.age)
		}
	})
}
