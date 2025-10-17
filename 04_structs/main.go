package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
	City string
}

// Method on struct
func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s, %d years old, from %s\n", p.Name, p.Age, p.City)
}

// Method with pointer receiver (can modify the struct)
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("%s is now %d years old!\n", p.Name, p.Age)
}

func main() {
	// Create struct instances
	person1 := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}

	person2 := Person{"Bob", 25, "London"}

	person1.Introduce()
	person2.Introduce()

	// Using pointer receiver method
	person1.HaveBirthday()
	person1.Introduce()
}
