# Go (Golang) Comprehensive Learning Guide

## 📌 Introduction
Go (Golang) is a modern, efficient, and statically typed programming language designed by Google. It is known for its simplicity, concurrency support, and high performance, making it ideal for backend development, microservices, system programming, and cloud-native applications.

---

## 📖 Learning Roadmap

### 1️⃣ Fundamentals
- **Setting Up Go**: Installing Go, configuring GOPATH, and using `go mod`
- **Basic Syntax**: Variables, constants, and data types
- **Control Structures**: Loops (`for`), conditionals (`if`, `switch`)
- **Type Conversion & Formatting**: Using `fmt.Printf` and type assertions
- **Shadowing & Scope**: Understanding variable scope and shadowing issues
- **Constants & Iota**: Declaring constants and using the `iota` enumerator

### 2️⃣ Functions & Error Handling
- **Functions**: Declaration, multiple return values, variadic functions
- **Error Handling**: `error` interface, `defer`, `panic`, and `recover`
- **Custom Errors**: Using `errors.New` and `fmt.Errorf`
- **Logging Errors**: Implementing structured logging with `log` and `zap`

### 3️⃣ Data Structures & Memory Management
- **Arrays, Slices, and Maps**
- **Structs, Methods, and Interfaces**
- **Pointers**: Memory allocation, `new()` vs `make()`, and dereferencing
- **Garbage Collection**: How Go manages memory automatically

### 4️⃣ Concurrency & Parallelism
- **Goroutines**: Lightweight threads for concurrent execution
- **Channels**: Communication between goroutines
- **Buffered vs Unbuffered Channels**: Understanding their differences
- **Select Statement**: Handling multiple channel operations
- **Sync Package**: Mutexes, WaitGroups, and atomic operations

### 5️⃣ Web Development & APIs
- **Building REST APIs**: Using frameworks like `Gin`, `Echo`, or `Fiber`
- **Middleware & Authentication**: JWT, OAuth, and role-based access
- **Working with Databases**: Using `GORM` and raw SQL queries
- **File Handling & JSON Parsing**
- **WebSockets & Real-Time Communication**

### 6️⃣ Microservices & Distributed Systems
- **gRPC & Protocol Buffers**
- **Message Queues**: Kafka, RabbitMQ integration
- **Service Discovery & Load Balancing**
- **Containerization with Docker & Kubernetes**

### 7️⃣ Testing & Debugging
- **Unit Testing**: Using `testing` package
- **Integration Testing**: API tests with `httptest`
- **Benchmarking & Profiling**: Optimizing performance with `pprof`
- **Local Testing Strategies**: Using `go test` with mocks

### 8️⃣ Deployment & DevOps
- **CI/CD Pipelines**: Automating Go builds and deployments
- **Cloud Deployment**: AWS, GCP, Azure
- **Logging & Monitoring**: Using `log`, `zap`, and Prometheus

### 9️⃣ Best Practices & Design Patterns
- **Clean Code Principles in Go**
- **SOLID Principles**
- **Common Go Design Patterns**: Factory, Singleton, Observer, etc.

---
## 🚀 Getting Started
### Install Go (Linux/Mac/Windows)
```sh
# Download and install Go from the official website
go version  # Verify installation
```
### Run Your First Go Program
```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```
Run the program:
```sh
go run main.go
```

---

## 🔗 Resources
- [Official Go Documentation](https://golang.org/doc/)
- [Go By Example](https://gobyexample.com/)
- [The Go Programming Language Book](https://www.gopl.io/)
- [Effective Go](https://golang.org/doc/effective_go.html)





























# Functions in Go: In-Depth Guide & Practice Problems

## Table of Contents
1. [Introduction to Functions](#introduction-to-functions)
2. [Defining Functions](#defining-functions)
3. [Function Parameters](#function-parameters)
4. [Return Values](#return-values)
5. [Multiple Return Values](#multiple-return-values)
6. [Named Return Values](#named-return-values)
7. [Variadic Functions](#variadic-functions)
8. [Anonymous Functions](#anonymous-functions)
9. [Higher-Order Functions](#higher-order-functions)
10. [Defer Keyword](#defer-keyword)
11. [Practice Problems](#practice-problems)
12. [Additional Practice Questions](#additional-practice-questions)

---

## Introduction to Functions
A **function** is a reusable block of code that performs a specific task. Functions promote:
- ✅ **Code Reusability**
- ✅ **Modularity**
- ✅ **Readability**
- ✅ **Maintainability**

---

## Defining Functions
In Go, use the `func` keyword to define a function.

### Example: Simple Function
```go
package main
import "fmt"

func greet() {
    fmt.Println("Hello, Welcome to Golang!")
}

func main() {
    greet() // Output: Hello, Welcome to Golang!
}
```

---

## Function Parameters
Functions can accept inputs (parameters).

### Example: Function with Parameters
```go
func add(a int, b int) {
    fmt.Println("Sum:", a + b)
}
```

---

## Return Values
Functions can return results.

### Example: Return a Value
```go
func add(a int, b int) int {
    return a + b
}
```

---

## Multiple Return Values
Go allows functions to return multiple values.

### Example: Return Quotient & Remainder
```go
func divide(dividend int, divisor int) (int, int) {
    return dividend / divisor, dividend % divisor
}
```

---

## Named Return Values
Predefine return variables for clarity.

### Example: Named Returns
```go
func rectangleProps(length, width int) (area int, perimeter int) {
    area = length * width
    perimeter = 2 * (length + width)
    return // Automatically returns area and perimeter
}
```

---

## Variadic Functions
Accept a variable number of arguments with `...`.

### Example: Sum of Numbers
```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
```

---

## Anonymous Functions
Functions without a name, defined inline.

### Example: Anonymous Function
```go
add := func(a, b int) int {
    return a + b
}
fmt.Println(add(5, 3)) // Output: 8
```

---

## Higher-Order Functions
Pass functions as arguments.

### Example: Function as Parameter
```go
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}
```

---

## Defer Keyword
Delay execution until the surrounding function returns.

### Example: Defer Usage
```go
func main() {
    defer fmt.Println("Deferred execution!")
    fmt.Println("Start") // Output: Start → End → Deferred execution!
}
```

---

## Practice Problems

### Basic Functions
1. Calculate the factorial of a number.
2. Check if a number is prime.
3. Count vowels in a string.

### Multiple Return Values
4. Return the min and max of an array.
5. Find the Nth Fibonacci number using recursion.

### Variadic Functions
6. Multiply any number of arguments.
7. Find the longest word in a sentence.

---

## Additional Practice Questions

### Beginner Level
1. Check if a number is even/odd.
2. Convert Celsius to Fahrenheit.
3. Reverse a string without built-in functions.

### Multiple Return Values
11. Return sum and product of two numbers.
12. Find smallest/largest number in an array.

### Recursion
16. Recursive factorial calculation.
17. Recursive Fibonacci sequence.

### Variadic Functions
21. Calculate average of integers.
22. Find maximum among given values.

### Higher-Order Functions
26. Apply mathematical operations dynamically.
29. Sort array using a callback function.
---




# 🚀 Go Closures: Stateful Functions Made Simple 🧩

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![Closures](https://img.shields.io/badge/Feature-Closures-success)

A comprehensive guide to mastering closures in Golang - with real-world examples and practical applications!

```go
package main

func main() {
    // The magic starts here ✨
}
```

---

## 🌟 Table of Contents

- [💡 What Are Closures?](#-what-are-closures)
- [🚀 Quick Start](#-quick-start)
- [🔍 Core Examples](#-core-examples)
  - [Basic Closure](#-basic-closure)
  - [Parameterized Closure](#-parameterized-closure)
  - [Real-World Use Case](#-real-world-api-rate-limiter)
- [✅ Key Benefits](#-key-benefits)
- [🧠 Why Use Closures?](#-why-use-closures)
- [📚 Learn More](#-learn-more)

---

## 💡 What Are Closures?

**Closures** are special functions that:
- 🔒 Capture variables from their surrounding scope
- 📦 Maintain state between calls
- 🎯 Enable functional programming patterns

> "A closure is a function value that references variables from outside its body" - [Go Language Spec](https://go.dev/ref/spec)

---

## 🚀 Quick Start

### Basic Counter Closure
```go
counter := func() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}()

fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
```

---

## 🔍 Core Examples

### 🎯 Basic Closure: Stateful Counter
```go
func main() {
    counter := createCounter()
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
}

func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```
**Key Insight:** The inner function maintains access to `count` even after `createCounter` exits

---

### 🎛️ Parameterized Closure: Multiplier Factory
```go
func multiplierFactory(n int) func(int) int {
    return func(x int) int {
        return x * n
    }
}

// Usage
double := multiplierFactory(2)
fmt.Println(double(5)) // 10
```

---

### 🌐 Real-World API Rate Limiter
```go
func rateLimiter(maxCalls int, duration time.Duration) func() bool {
    count := 0
    lastReset := time.Now()
    
    return func() bool {
        if time.Since(lastReset) > duration {
            count = 0
            lastReset = time.Now()
        }
        if count < maxCalls {
            count++
            return true
        }
        return false
    }
}

// Usage: Allow 3 requests/second
checkLimit := rateLimiter(3, time.Second)
```

---

## ✅ Key Benefits

| Benefit | Description |
|---------|-------------|
| 🧠 **State Persistence** | Maintain values between function calls |
| 🛡️ **Encapsulation** | Keep variables private to the closure |
| 🔧 **Flexibility** | Create dynamic function factories |
| 🚀 **Performance** | Lightweight alternative to structs |

---

## 🧠 Why Use Closures?

1. **State Management**  
   Perfect for counters, rate limiters, generators

2. **Function Factories**  
   Create specialized functions on-the-fly

3. **Callback Handlers**  
   Maintain context in async operations

4. **Testing**  
   Mock external dependencies easily

---

## 📚 Learn More

### Recommended Resources
- [Go Tour: Closures](https://go.dev/tour/moretypes/25)
- [Effective Go: Functions](https://go.dev/doc/effective_go#functions)
- [Closures vs Structs](https://www.calhoun.io/using-closures-in-go/) (Performance comparison)

---

<!-- <div align="center">
  <h3>💬 Found a Bug? Have a Suggestion?</h3>
  <p>We love contributions!<br>
  <a href="https://github.com/your-repo/issues">Open an Issue</a> or 
  <a href="https://github.com/your-repo/pulls">Submit a PR</a></p>
  
  [![Go Report Card](https://goreportcard.com/badge/github.com/your-repo)](https://goreportcard.com/report/github.com/your-repo)
</div> -->



# Error Handling in Go: A Deep Dive

Go follows an explicit error-handling approach rather than relying on exceptions like many other languages (e.g., Python, Java). In Go, errors are values, which makes handling them predictable and clear.

## What is an Error in Go?

In Go, an error is simply a value of type `error`. Functions that may fail usually return an error type as their last return value. The caller must explicitly check and handle errors.

## Basic Syntax of Error Handling

```go
package main

import (
	"fmt"
	"errors"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
```

### Key Points:
- If the divisor is zero, return an error using `errors.New()`.
- If there's no error, return `nil` as the second value.
- The caller checks if `err != nil` and handles it accordingly.

## Creating Custom Errors

Go allows custom errors to provide more detailed error messages.

### Using `fmt.Errorf()`

```go
package main

import (
	"fmt"
)

func getUser (id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("invalid user ID: %d", id)
	}
	return "User 123", nil
}

func main() {
	user, err := getUser (-1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User :", user)
	}
}
```

### Key Point:
- `fmt.Errorf()` allows formatted error messages with values.

## Using `errors.Is()` and `errors.As()` for Error Wrapping

Go 1.13 introduced error wrapping, making error comparison easier.

### Example: Wrapping and Checking Errors

```go
package main

import (
	"fmt"
	"errors"
)

var ErrNotFound = errors.New("item not found")

func findItem(id int) error {
	if id != 1 {
		return fmt.Errorf("findItem: %w", ErrNotFound) // Wrapping error
	}
	return nil
}

func main() {
	err := findItem(2)
	if errors.Is(err, ErrNotFound) { // Checking error type
		fmt.Println("Custom Error: Item not found")
	} else if err != nil {
		fmt.Println("Error:", err)
	}
}
```

### Key Point:
- `errors.Is()` checks if an error matches a specific type, even if wrapped.

## Using Panic and Recover

### What is Panic?

`panic` stops the normal execution of a program. It should only be used for unrecoverable errors (e.g., array index out of bounds).

### Example: Using Panic

```go
package main

import "fmt"

func main() {
	fmt.Println("Before panic")
	panic("Something went wrong!") // This will stop execution
	fmt.Println("After panic") // This line won't execute
}
```

### Using Recover to Handle Panic

`recover()` is used inside `defer` to catch and handle a panic gracefully. It prevents the program from crashing.

```go
package main

import "fmt"

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println(a / b) // This will panic if b = 0
}

func main() {
	safeDivide(10, 0)  // No crash, thanks to `recover`
	fmt.Println("Program continues...")
}
```

### Key Point:
- `recover()` captures the panic and prevents the program from crashing.

## Best Practices for Error Handling in Go

- Always check errors when calling functions that return error.
- Return errors instead of panicking, unless it's truly unrecoverable.
- Wrap errors using `fmt.Errorf()` for better debugging.
- Use `errors.Is()` and `errors.As()` for comparing errors properly.
- Use `defer` and `recover` only for critical failures (e.g., handling server crashes).

## Error Handling Practice Questions

1. Write a function to open a file and return an error if it doesn't exist.
2. Implement a function that parses an integer from a string and handles conversion errors.
3. Create a function that reads user input and returns an error if it's empty.
4. Write a function that checks an array index and prevents out-of-bounds errors.
5. Implement a function that calculates the square root of a number, returning an error if the number is negative.

---

