package mathutil

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zeros", 0, 0, 0},
		{"large numbers", 1000000, 2000000, 3000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive result", 5, 3, 2},
		{"negative result", 3, 5, -2},
		{"zeros", 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 3, 4, 12},
		{"by zero", 5, 0, 0},
		{"negative numbers", -3, -4, 12},
		{"mixed signs", -3, 4, -12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("valid division", func(t *testing.T) {
		result, err := Divide(10, 3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if math.Abs(result-3.3333333333333335) > 1e-9 {
			t.Errorf("Divide(10, 3) = %f; want ~3.333", result)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Fatal("expected error for division by zero")
		}
	})
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		expected  int
		expectErr bool
	}{
		{"zero", 0, 1, false},
		{"one", 1, 1, false},
		{"five", 5, 120, false},
		{"ten", 10, 3628800, false},
		{"negative", -1, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Factorial(tt.n)
			if tt.expectErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Factorial(%d) = %d; want %d", tt.n, result, tt.expected)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"two", 2, true},
		{"three", 3, true},
		{"four", 4, false},
		{"seventeen", 17, true},
		{"twenty", 20, false},
		{"ninety-seven", 97, true},
		{"negative", -5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPrime(tt.n)
			if result != tt.expected {
				t.Errorf("IsPrime(%d) = %v; want %v", tt.n, result, tt.expected)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		expected  int
		expectErr bool
	}{
		{"zero", 0, 0, false},
		{"one", 1, 1, false},
		{"six", 6, 8, false},
		{"ten", 10, 55, false},
		{"negative", -1, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Fibonacci(tt.n)
			if tt.expectErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Fibonacci(%d) = %d; want %d", tt.n, result, tt.expected)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"basic", 12, 8, 4},
		{"coprime", 7, 13, 1},
		{"same number", 5, 5, 5},
		{"one is zero", 0, 5, 5},
		{"both zero", 0, 0, 0},
		{"negative inputs", -12, 8, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GCD(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("GCD(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
