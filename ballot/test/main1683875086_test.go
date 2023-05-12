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

func (p *Product) setId(id int) {
	p.id = id
}

func TestSetId(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "Positive ID",
			input: 1,
			want:  1,
		},
		{
			name:  "Zero ID",
			input: 0,
			want:  0,
		},
		{
			name:  "Negative ID",
			input: -1,
			want:  -1,
		},
		{
			name:  "Max Int ID",
			input: 1<<31 - 1,
			want:  1<<31 - 1,
		},
		{
			name:  "Min Int ID",
			input: -1 << 31,
			want:  -1 << 31,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			product := &Product{}
			product.setId(test.input)
			if product.id != test.want {
				t.Errorf("setId(%d) = %d; want %d", test.input, product.id, test.want)
			}
		})
	}
}

