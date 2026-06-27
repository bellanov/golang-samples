package calculator

import (
	"testing"
)

// TestAdd tests the Add method.
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed signs", 10, -5, 5},
		{"zero", 0, 0, 0},
	}

	calc := NewCalculator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestSubtract tests the Subtract method.
func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 10, 3, 7},
		{"negative numbers", -10, -3, -7},
		{"mixed signs", 10, -5, 15},
		{"zero", 0, 0, 0},
	}

	calc := NewCalculator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestMultiply tests the Multiply method.
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 6, 7, 42},
		{"negative numbers", -6, -7, 42},
		{"mixed signs", 6, -7, -42},
		{"zero", 5, 0, 0},
	}

	calc := NewCalculator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestDivide tests the Divide method.
func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		expected  int
		wantError bool
	}{
		{"positive numbers", 20, 4, 5, false},
		{"negative numbers", -20, -4, 5, false},
		{"mixed signs", 20, -4, -5, false},
		{"division by zero", 20, 0, 0, true},
		{"zero divided", 0, 5, 0, false},
	}

	calc := NewCalculator()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, tt.b)
			if (err != nil) != tt.wantError {
				t.Errorf("Divide(%d, %d) error = %v; wantError %v", tt.a, tt.b, err, tt.wantError)
			}
			if !tt.wantError && result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
