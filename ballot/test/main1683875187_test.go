package main

import (
	"testing"
)

type Product struct {
	name     string
	category string
	price    float64
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
		{"Empty name", "", ""},
		{"Valid name", "Product A", "Product A"},
		{"Long name", "A very long product name that exceeds the limit", "A very long product name that exceeds the limit"},
		{"Special characters", "@#$%^&*()", "@#$%^&*()"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			product := &Product{}
			product.setName(test.input)
			if product.name != test.expected {
				t.Errorf("Expected name: %s, got: %s", test.expected, product.name)
			}
		})
	}
}
