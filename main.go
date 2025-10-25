package main

import (
    "fmt"
)

// cukicahamas-16 - Go Implementation
// A simple Hello World program

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Welcome to cukicahamas-16")
    
    // Simple calculation
    numbers := []int{1, 2, 3, 4, 5}
    total := 0
    for _, num := range numbers {
        total += num
    }
    fmt.Printf("Sum of numbers: %d\n", total)
}
