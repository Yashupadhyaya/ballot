package main

import (
	"testing"
)

type Product struct {
	id       int
	name     string
	category string
	price    float64
}

func (p *Product) getId() int {
	return p.id
}

func TestGetId(t *testing.T) {
	tests := []struct {
		name     string
		product  Product
		expected int
	}{
		{
			name:     "Positive ID",
			product:  Product{id: 1, name: "Product 1", category: "Electronics", price: 100.0},
			expected: 1,
		},
		{
			name:     "Zero ID",
			product:  Product{id: 0, name: "Product 2", category: "Electronics", price: 200.0},
			expected: 0,
		},
		{
			name:     "Negative ID",
			product:  Product{id: -1, name: "Product 3", category: "Electronics", price: 300.0},
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.product.getId()
			if result != test.expected {
				t.Errorf("Expected ID %d, but got %d", test.expected, result)
			}
		})
	}
}
