package main

import (
	"errors"
	"fmt"
	"os"
)

// Custom error type
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Msg)
}

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function with custom error
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Msg: "must be realistic"}
	}
	return nil
}

// Function that demonstrates panic and recover
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero!")
	}
	result = a / b
	return result, nil
}

func main() {
	fmt.Println("=== Basic Error Handling ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n=== Custom Error ===")
	if err := validateAge(30); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid")
	}

	if err := validateAge(-5); err != nil {
		fmt.Println("Error:", err)
		if valErr, ok := err.(*ValidationError); ok {
			fmt.Printf("Field: %s\n", valErr.Field)
		}
	}

	fmt.Println("\n=== Panic and Recover ===")
	intResult, err := safeDivide(10, 2)
	fmt.Printf("Result: %d, Error: %v\n", intResult, err)

	intResult, err = safeDivide(10, 0)
	fmt.Printf("Result: %d, Error: %v\n", intResult, err)

	fmt.Println("\n=== Error Wrapping (Go 1.13+) ===")
	_, err = os.Open("/nonexistent/file.txt")
	if err != nil {
		wrappedErr := fmt.Errorf("failed to open config: %w", err)
		fmt.Println("Wrapped error:", wrappedErr)

		// Unwrap to check the original error
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")
		}
	}
}
