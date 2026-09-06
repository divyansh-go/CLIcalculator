package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("\n===========================================")
	fmt.Println("              GO CALCULATOR")
	fmt.Println("===========================================")

	resp := "y"
	for resp == "y" || resp == "Y" {

		cal()

		fmt.Print("\nDo you want to perform another calculation? (y/n): ")
		fmt.Scan(&resp)

		for resp != "y" && resp != "Y" && resp != "n" && resp != "N" {

			fmt.Println("\nERROR: Invalid Input. Try Again......")

			fmt.Print("Do you want to perform another calculation? (y/n): ")
			fmt.Scan(&resp)

		}

	}

	fmt.Println("\nGoodbye!")
}

func cal() {

	var n1, n2, r float64
	var op string

	fmt.Print("\nEnter first number: ")
	_, err := fmt.Scan(&n1)
	val := errhandle(err)
	if val == 1 {
		return
	}

	fmt.Print("Enter operation (+, -, *, /): ")
	fmt.Scan(&op)
	switch op {
	case "+", "-", "*", "/":
	default:
		fmt.Println("\nERROR: Invalid Operation. Please choose +, -, *, or /.")
		return
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&n2)
	val = errhandle(err)
	if val == 1 {
		return
	}

	switch op {
	case "+":
		r = n1 + n2
	case "-":
		r = n1 - n2
	case "*":
		r = n1 * n2
	case "/":
		if n2 == 0 {
			fmt.Println("\nERROR: cannot divide by zero.")
			return
		} else {
			r = n1 / n2
		}
	}

	fmt.Println("\nResult: ", r)
}

func errhandle(err1 error) int {
	reader := bufio.NewReader(os.Stdin)
	if err1 != nil {
		fmt.Println("\nERROR: Invalid Number. Please enter a valid number.")
		reader.ReadString('\n')
		return 1
	}
	return 0
}
