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
	t.Run("Test positive age", func(t *testing.T) {
		p := &Person{}
		p.setAge(25)
		if p.age != 25 {
			t.Errorf("Expected age to be 25, got %d", p.age)
		}
	})

	t.Run("Test negative age", func(t *testing.T) {
		p := &Person{}
		p.setAge(-5)
		if p.age != 0 {
			t.Errorf("Expected age to be 0, got %d", p.age)
		}
	})

	t.Run("Test zero age", func(t *testing.T) {
		p := &Person{}
		p.setAge(0)
		if p.age != 0 {
			t.Errorf("Expected age to be 0, got %d", p.age)
		}
	})

	t.Run("Test maximum age", func(t *testing.T) {
		p := &Person{}
		p.setAge(150)
		if p.age != 150 {
			t.Errorf("Expected age to be 150, got %d", p.age)
		}
	})

	t.Run("Test age beyond maximum", func(t *testing.T) {
		p := &Person{}
		p.setAge(200)
		if p.age != 0 {
			t.Errorf("Expected age to be 0, got %d", p.age)
		}
	})
}
