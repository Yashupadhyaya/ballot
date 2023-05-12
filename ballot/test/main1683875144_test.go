package main

import (
	"testing"
)

type Product struct {
	name     string
	category string
	price    float64
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
			name:     "Success: Get product name",
			product:  Product{name: "Product A", category: "Electronics", price: 100.0},
			expected: "Product A",
		},
		{
			name:     "Success: Get empty product name",
			product:  Product{name: "", category: "Electronics", price: 100.0},
			expected: "",
		},
		{
			name:     "Success: Get product name with special characters",
			product:  Product{name: "@#$%^&*()", category: "Electronics", price: 100.0},
			expected: "@#$%^&*()",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.product.getName()
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}
