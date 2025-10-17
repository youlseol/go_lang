package main

import "fmt"

// Simple function
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Function with return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Named return values
func rectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

func main() {
	greet("Gopher")

	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", area, perimeter)
}
