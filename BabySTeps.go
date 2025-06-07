package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This is a simple calculator program that performs basic arithmetic operations.
// It prompts the user for two numbers and an operator, performs the operation,
// and displays the result.

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

// divide performs integer division. Consider float64 for decimal results.
func divide(a, b int) int {
	return a / b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the first number: ")
	scanner.Scan()
	input1 := scanner.Text()
	num1, err := strconv.Atoi(strings.TrimSpace(input1))
	if err != nil {
		fmt.Println("nah bra. try a valid input.")
		return
	}

	fmt.Print("Enter the second number: ")
	scanner.Scan()
	input2 := scanner.Text()
	num2, err := strconv.Atoi(strings.TrimSpace(input2))
	if err != nil {
		fmt.Println("nah bra. try a valid input.")
		return
	}

	fmt.Print("Enter an operator (+, -, *, /): ")
	scanner.Scan()
	operator := strings.TrimSpace(scanner.Text())

	switch operator {
	case "+":
		fmt.Printf("Result: %d\n", add(num1, num2))
	case "-":
		fmt.Printf("Result: %d\n", subtract(num1, num2))
	case "*":
		fmt.Printf("Result: %d\n", multiply(num1, num2))
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero r u dumb.") // Keeping the original tone
			return
		}
		fmt.Printf("Result: %d\n", divide(num1, num2))
	default:
		fmt.Println("Stick to the script bud. Please enter one of +, -, *, /.")
	}

	fmt.Println("Thank you for using BabySTeps by xa tamba !")
}
