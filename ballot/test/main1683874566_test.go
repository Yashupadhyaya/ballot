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
	{3, 19, "Charlie"},
}

func listStudent() {
	for _, student := range studentList {
		println(student.id, student.age, student.name)
	}
}

func TestListStudent(t *testing.T) {
	// Test case 1: Check if the listStudent function prints the correct output
	expectedOutput := "1 20 Alice\n2 22 Bob\n3 19 Charlie\n"
	output := captureOutput(listStudent)

	if output != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output)
	}

	// Test case 2: Check if the listStudent function handles an empty list
	studentList = []Student{}
	expectedOutput = ""
	output = captureOutput(listStudent)

	if output != expectedOutput {
		t.Errorf("Expected output: %s, got: %s", expectedOutput, output)
	}
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
