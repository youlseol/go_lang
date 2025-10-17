package main

import (
	"testing"
)

func TestValidateAge(t *testing.T) {
	tests := []struct {
		name      string
		age       int
		expectErr bool
	}{
		{"valid age", 30, false},
		{"negative age", -5, true},
		{"unrealistic age", 200, true},
		{"zero age", 0, false},
		{"boundary age", 150, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAge(tt.age)
			if tt.expectErr && err == nil {
				t.Errorf("validateAge(%d) expected error, got nil", tt.age)
			}
			if !tt.expectErr && err != nil {
				t.Errorf("validateAge(%d) unexpected error: %v", tt.age, err)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("divide(10, 2) unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("divide(10, 2) = %f; want 5.0", result)
	}

	_, err = divide(10, 0)
	if err == nil {
		t.Error("divide(10, 0) expected error, got nil")
	}
}

func TestSafeDivide(t *testing.T) {
	result, err := safeDivide(10, 2)
	if err != nil {
		t.Errorf("safeDivide(10, 2) unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("safeDivide(10, 2) = %d; want 5", result)
	}

	result, err = safeDivide(10, 0)
	if err == nil {
		t.Error("safeDivide(10, 0) expected error from panic, got nil")
	}
	if result != 0 {
		t.Errorf("safeDivide(10, 0) = %d; want 0", result)
	}
}

func TestValidationError(t *testing.T) {
	err := &ValidationError{Field: "test", Msg: "test message"}
	expected := "validation error on field 'test': test message"
	if err.Error() != expected {
		t.Errorf("ValidationError.Error() = %s; want %s", err.Error(), expected)
	}
}
