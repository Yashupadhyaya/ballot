package main

import (
	"testing"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

type App struct {
	studentList []Student
}

func (app *App) addStudent(student Student) {
	app.studentList = append(app.studentList, student)
}

func (app *App) updateStudent(id int, updatedStudent Student) {
	for i, student := range app.studentList {
		if student.ID == id {
			app.studentList[i] = updatedStudent
			break
		}
	}
}

func (app *App) deleteStudent(id int) {
	for i, student := range app.studentList {
		if student.ID == id {
			app.studentList = append(app.studentList[:i], app.studentList[i+1:]...)
			break
		}
	}
}

func (app *App) fetchStudent(id int) *Student {
	for _, student := range app.studentList {
		if student.ID == id {
			return &student
		}
	}
	return nil
}

func (app *App) isPresent(id int) bool {
	for _, student := range app.studentList {
		if student.ID == id {
			return true
		}
	}
	return false
}

func (app *App) listStudent() []Student {
	return app.studentList
}

func TestApp(t *testing.T) {
	app := App{}

	// Test addStudent
	student1 := Student{ID: 1, Name: "John", Age: 20}
	app.addStudent(student1)
	if len(app.studentList) != 1 {
		t.Errorf("addStudent failed, expected length 1, got %d", len(app.studentList))
	}

	// Test updateStudent
	updatedStudent := Student{ID: 1, Name: "John Doe", Age: 21}
	app.updateStudent(1, updatedStudent)
	if app.studentList[0].Name != "John Doe" || app.studentList[0].Age != 21 {
		t.Errorf("updateStudent failed, expected updated student, got %+v", app.studentList[0])
	}

	// Test deleteStudent
	app.deleteStudent(1)
	if len(app.studentList) != 0 {
		t.Errorf("deleteStudent failed, expected length 0, got %d", len(app.studentList))
	}

	// Test fetchStudent
	app.addStudent(student1)
	fetchedStudent := app.fetchStudent(1)
	if fetchedStudent == nil || fetchedStudent.ID != 1 || fetchedStudent.Name != "John" || fetchedStudent.Age != 20 {
		t.Errorf("fetchStudent failed, expected student1, got %+v", fetchedStudent)
	}

	// Test isPresent
	if !app.isPresent(1) {
		t.Errorf("isPresent failed, expected true, got false")
	}

	// Test listStudent
	studentList := app.listStudent()
	if len(studentList) != 1 || studentList[0].ID != 1 || studentList[0].Name != "John" || studentList[0].Age != 20 {
		t.Errorf("listStudent failed, expected studentList with student1, got %+v", studentList)
	}
}
