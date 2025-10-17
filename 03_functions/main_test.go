package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(5, 3)
	expected := 8
	if result != expected {
		t.Errorf("add(5, 3) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		expectErr bool
	}{
		{"valid division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative numbers", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := divide(tt.a, tt.b)
			if tt.expectErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("divide(%f, %f) = %f; want %f", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

func TestRectangle(t *testing.T) {
	area, perimeter := rectangle(5, 3)
	expectedArea := 15.0
	expectedPerimeter := 16.0

	if area != expectedArea {
		t.Errorf("rectangle area = %f; want %f", area, expectedArea)
	}
	if perimeter != expectedPerimeter {
		t.Errorf("rectangle perimeter = %f; want %f", perimeter, expectedPerimeter)
	}
}
