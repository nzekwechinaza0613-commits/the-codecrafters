package main

import (
	"fmt"
	"strconv"
)

func style(text, style, color string) string {
	return style + color + text + "\033[0m"
}
func main() {

	fmt.Println(" ")
	fmt.Println(style("*************** WELCOME TO BLASIE PASCAL'S CALCULATOR ****************", "\033[1;33m", ""))
	fmt.Println(" ")
	fmt.Println(style("PRESS <help> TO SEE AVAILABLE OPERATIONS", "\033[1;33m", ""))
	fmt.Println(" ")

start:
	fmt.Print("FIRST VALUE: ")

	var first_value string
	fmt.Scanln(&first_value)

	switch first_value {

	case "quit", "QUIT":
		fmt.Println(style("GOODBYE👋!!", "\033[1;32m", ""))
		return

	case "help", "HELP":
		fmt.Println(style("-------- AVAILABLE COMMAND -----------", "\033[1;33m", ""))
		fmt.Println(style("<num1> + <num2>   → addition", "\033[1;30m", ""))
		fmt.Println(style("<num1> - <num2>   → subtraction", "\033[1;30m", ""))
		fmt.Println(style("<num1> * <num2>   → multiplication", "\033[1;30m", ""))
		fmt.Println(style("<num1> / <num2>   → division", "\033[1;30m", ""))
		fmt.Println(style("<quit>   → exit", "\033[1;30m", ""))
		fmt.Println(" ")
		goto start

	}

	num1, err := strconv.ParseFloat(first_value, 64)

	if err != nil {
		fmt.Println(style("VALUE MUST BE AN INTEGER", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println(" ")
		goto start

	}

operate:

	fmt.Println(" ")
	fmt.Println(style("AVAILABLE OPERATORS: add(+), subtract(-), multiply(*), divide(/)", "\033[1;36m", ""))
	fmt.Println(" ")
	fmt.Print("OPERATOR: ")
	var sign string
	fmt.Scanln(&sign)

	if sign == "quit" || sign == "QUIT" {
		fmt.Println(style("GOODBYE👋!!", "\033[1;32m", ""))
		return
	}

second:
	fmt.Println(" ")
	fmt.Print("SECOND VALUE: ")
	var second_value string
	fmt.Scanln(&second_value)

	switch second_value {
	case "quit", "QUIT":
		fmt.Println(style("GOODBYE👋!!", "\033[1;32m", ""))
		return

	case "help", "HELP":
		fmt.Println(style("-------- AVAILABLE COMMAND -----------", "\033[1;33m", ""))
		fmt.Println(style("<num1> + <num2>   → addition", "\033[1;30m", ""))
		fmt.Println(style("<num1> - <num2>   → subtraction", "\033[1;30m", ""))
		fmt.Println(style("<num1> * <num2>   → multiplication", "\033[1;30m", ""))
		fmt.Println(style("<num1> / <num2>   → division", "\033[1;30m", ""))
		fmt.Println(style("<quit>   → exit", "\033[1;30m", ""))
		fmt.Println(" ")
		goto second

	}

	num2, err := strconv.ParseFloat(second_value, 64)

	if err != nil {
		fmt.Println(style("VALUE MUST BE AN INTEGER", "\033[1;31m", ""))
		fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
		fmt.Println(" ")
		goto second
	}

	switch sign {

	case "+", "ADD", "add":
		fmt.Printf(style("RESULT: %.2f %s %.2f = %.2f\n", "\033[1;32m", ""), num1, sign, num2, num1+num2)

	case "-", "sub", "SUB":
		fmt.Printf(style("RESULT: %.2f %s %.2f = %.2f\n", "\033[1;32m", ""), num1, sign, num2, num1-num2)

	case "*", "MUL", "mul":
		fmt.Printf(style("RESULT: %.2f %s %.2f = %.2f\n", "\033[1;32m", ""), num1, sign, num2, num1*num2)

	case "/", "div", "DIV":
		if second_value == "0" {
			fmt.Println(style("THE DIVISOR CANNOT BE ZERO", "\033[1;31m", ""))
			fmt.Println(style("TRY AGAIN", "\033[1;31m", ""))
			fmt.Println(" ")
			goto second

		}
		fmt.Printf(style("RESULT: %.2f %s %.2f = %.2f\n", "\033[1;32m", ""), num1, sign, num2, num1/num2)

	case "help", "HELP":
		fmt.Println(style("-------- AVAILABLE COMMAND -----------", "\033[1;33m", ""))
		fmt.Println(style("<num1> + <num2>   → addition", "\033[1;30m", ""))
		fmt.Println(style("<num1> - <num2>   → subtraction", "\033[1;30m", ""))
		fmt.Println(style("<num1> * <num2>   → multiplication", "\033[1;30m", ""))
		fmt.Println(style("<num1> / <num2>   → division", "\033[1;30m", ""))
		fmt.Println(style("<quit>   → exit", "\033[1;30m", ""))
		fmt.Println(" ")
		goto operate

	default:
		fmt.Println(style("INVALID ARITHMETIC SYNTAX", "\033[1;31m", ""))
		fmt.Println(" ")
		fmt.Println(style("PRESS <help> TO SEE AVAILABLE OPERATIONS", "\033[1;33m", ""))
		fmt.Println(" ")
		goto start
	}
	fmt.Println(" ")
	goto start

}
