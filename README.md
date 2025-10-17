# Go Language Learning Repository

Welcome to the Go (Golang) learning repository! This repository contains practical examples to help you learn Go programming from basics to advanced concepts.

## 📚 Table of Contents

1. [Hello World](#01-hello-world) - Your first Go program
2. [Variables](#02-variables) - Variable declarations and constants
3. [Functions](#03-functions) - Function definitions, parameters, and return values
4. [Structs](#04-structs) - Custom data types and methods
5. [Interfaces](#05-interfaces) - Polymorphism in Go
6. [Concurrency](#06-concurrency) - Goroutines, channels, and WaitGroups
7. [Error Handling](#07-error-handling) - Error handling patterns and best practices

## 🚀 Getting Started

### Prerequisites

- Go 1.16 or higher installed on your system
- Basic understanding of programming concepts

### Installation

1. Clone this repository:
```bash
git clone https://github.com/youlseol/go_lang.git
cd go_lang
```

2. Initialize the Go module (if not already done):
```bash
go mod init github.com/youlseol/go_lang
```

## 📖 Learning Path

### 01. Hello World

The classic first program - learn about packages, imports, and the main function.

```bash
cd 01_hello_world
go run main.go
```

**Key Concepts:**
- Package declaration
- Importing packages
- The `main` function
- `fmt.Println()` for output

### 02. Variables

Learn different ways to declare variables and constants in Go.

```bash
cd 02_variables
go run main.go
```

**Key Concepts:**
- Variable declarations with `var`
- Short variable declaration with `:=`
- Type inference
- Constants with `const`

### 03. Functions

Understand how to write and use functions in Go.

```bash
cd 03_functions
go run main.go
```

**Key Concepts:**
- Function declarations
- Parameters and return values
- Multiple return values
- Named return values
- Error handling in functions

### 04. Structs

Learn about custom data types and object-oriented programming in Go.

```bash
cd 04_structs
go run main.go
```

**Key Concepts:**
- Struct definitions
- Creating struct instances
- Methods on structs
- Pointer receivers vs value receivers

### 05. Interfaces

Master polymorphism and abstraction using interfaces.

```bash
cd 05_interfaces
go run main.go
```

**Key Concepts:**
- Interface definitions
- Implementing interfaces
- Polymorphism
- Type assertions

### 06. Concurrency

Explore Go's powerful concurrency features.

```bash
cd 06_concurrency
go run main.go
```

**Key Concepts:**
- Goroutines
- Channels
- WaitGroups
- Worker pools
- Concurrent programming patterns

### 07. Error Handling

Learn Go's idiomatic error handling patterns.

```bash
cd 07_error_handling
go run main.go
```

**Key Concepts:**
- Error interface
- Custom error types
- Error wrapping (Go 1.13+)
- Panic and recover
- Best practices for error handling

## 🧪 Running Tests

Each example can be run individually using:

```bash
cd <example_directory>
go run main.go
```

To run all examples:

```bash
for dir in 0*; do
    echo "Running $dir..."
    (cd "$dir" && go run main.go)
    echo ""
done
```

## 📝 Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Tour](https://tour.golang.org/)

## 🤝 Contributing

Feel free to contribute by:
- Adding more examples
- Improving documentation
- Fixing bugs
- Suggesting improvements

## 📄 License

This repository is for educational purposes. Feel free to use and modify the code as needed for your learning journey.

## 🎯 Next Steps

After completing these examples, consider:
- Building a small CLI application
- Creating a REST API with Go
- Exploring Go's standard library packages
- Learning about Go modules and dependency management
- Diving into testing with Go's `testing` package

Happy coding! 🎉