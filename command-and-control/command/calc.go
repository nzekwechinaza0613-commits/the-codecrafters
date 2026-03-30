package command

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculator(a, b string) {
	reader := bufio.NewReader(os.Stdin)

start:
	fmt.Print("C&C>")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.ToLower(input) == "quit" {
		fmt.Print(" Shutting down String Transformer.\n Goodbye.")
		return
	}

	num1, err := strconv.ParseFloat(a, 64)

	if err != nil {
		fmt.Println("VALUE MUST BE AN INTEGER")
		fmt.Println("TRY AGAIN")
		fmt.Println(" ")
		goto start
	}

	num2, err2 := strconv.ParseFloat(b, 64)

	if err2 != nil {
		fmt.Println("VALUE MUST BE AN INTEGER")
		fmt.Println("TRY AGAIN")
		fmt.Println(" ")
		goto start
	}

	parts := strings.Fields(input)
	command := parts[0]

	textSlice := parts[1:]

	text := strings.Join(textSlice, " ")

	if len(text) <= 2 {
		fmt.Println(" ")
		fmt.Println("USAGE: <command> <a> <b>")
		fmt.Println(" ")
	}

	switch strings.ToLower(command) {
	case "add":
		fmt.Printf("RESULT: %.2f\n", num1+num2)

	case "sub":
		fmt.Printf("RESULT: %.2f\n", num1-num2)

	case "mul":
		fmt.Printf("RESULT:  %.2f\n", num1*num2)

	case "div":
		if text == "0" {
			fmt.Println("THE DIVISOR CANNOT BE ZERO")
			fmt.Println("TRY AGAIN")
			fmt.Println(" ")
			goto start

		}
		fmt.Printf("RESULT: %.2f\n", num1/num2)

	// case "help", "HELP":
	// 	fmt.Println("-------- AVAILABLE COMMAND -----------")
	// 	fmt.Println("<num1> + <num2>   → addition")
	// 	fmt.Println("<num1> - <num2>   → subtraction")
	// 	fmt.Println("<num1> * <num2>   → multiplication")
	// 	fmt.Println("<num1> / <num2>   → division")
	// 	fmt.Println("<quit>   → exit")
	// 	fmt.Println(" ")

	case "mod":
		fmt.Printf("RESULT:  %v\n", int(num1)%int(num2))

	case "pow":
		fmt.Printf("RESULT:  %v\n", int(num1)^int(num2))

	default:
		fmt.Println("INVALID ARITHMETIC SYNTAX")
		fmt.Println(" ")
		goto start
	}

}
