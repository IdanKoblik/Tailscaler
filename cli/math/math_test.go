package math

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		y        int
		expected int
	}{
		{"Positive numbers", 3, 7, 7},
		{"Negative numbers", -5, -3, -3},
		{"Positive and negative numbers", 5, -2, 5},
		{"Equal numbers", 4, 4, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Max(test.x, test.y)
			if got != test.expected {
				t.Errorf("Max(%d, %d) = %d; expected %d", test.x, test.y, got, test.expected)
			}
		})
	}
}
