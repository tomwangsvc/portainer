package main

import (
	"testing"
)

func Test_Valid(t *testing.T) {
	container := NewContainer([][2]string{{"{", "}"}, {"[", "]"}, {"(", ")"}})
	var data = []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			input:    "",
			expected: true,
		},
		{
			input:    "{}",
			expected: true,
		},
		{
			input:    "[]",
			expected: true,
		},
		{
			input:    "{[]}",
			expected: true,
		},
		{
			input:    "{[{[]}]}",
			expected: true,
		},
		{
			input:    "{}[]()(([]))",
			expected: true,
		},
		{
			input:    "{}}",
			expected: false,
		},
		{
			input:    "{",
			expected: false,
		},
		{
			input:    "{[[]}",
			expected: false,
		},
		{
			input:    "{[}",
			expected: false,
		},
	}

	for i, d := range data {
		t.Run(d.desc, func(t *testing.T) {
			result := container.Valid(d.input)
			if result != d.expected {
				t.Errorf("result not match at: %d,  input: %+v, expected: %+v, result: %+v", i, d.input, d.expected, result)
			}
		})
	}
}
