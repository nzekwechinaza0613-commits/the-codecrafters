package main

import (
	"fmt"
	"strconv"
	"strings"
)

func color(text, color string) string {
	return color + text + "\033[0m"
}

func main() {

	fmt.Println(color("*************** WELCOME TO WILHEM SCHIKRED CONVERTER ****************", "\033[1;33m"))
	fmt.Println(" ")

start:
	fmt.Print("INPUT: ")
	var value string
	fmt.Scanln(&value)

	if value == "quit" || value == "QUIT" {
		fmt.Println(" ")
		fmt.Println(color("GOODBYE👋!!", "\033[1;32m"))
		return
	}

operate:
	fmt.Println(" ")
	fmt.Println("AVAILABLE CONVERSION: (<hex>, <bin>, <dec>)")
	fmt.Println(" ")
	fmt.Print("WHICH CONVERSION WOULD YOU WANT TO APPLY: ")
	var convertor string
	fmt.Scanln(&convertor)

	switch convertor {

	case "hex", "HEX":
		val, err := strconv.ParseInt(value, 16, 64)

		if err != nil {
			fmt.Println("VALUE IS NOT A HEXADECIMAL NUMBER")
			fmt.Println("TRY AGAIN")
			fmt.Println(" ")
			goto start
		}
		hex := strings.ToUpper(strconv.FormatInt(val, 10))
		fmt.Printf("DECIMAL = %v\n", hex)

	case "bin", "BIN":
		val, err := strconv.ParseInt(value, 2, 64)

		if err != nil {
			fmt.Println("VALUE IS NOT A BINARY NUMBER")
			fmt.Println("TRY AGAIN")
			fmt.Println(" ")
			goto start
		}
		fmt.Printf("DECIMAL = %v\n", val)

	case "dec", "DEC":
		val, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			fmt.Println("VALUE IS NOT A DECIMAL NUMBER")
			fmt.Println("TRY AGAIN")
			fmt.Println(" ")
			goto start
		}

		binary := strconv.FormatInt(val, 2)
		hex := strings.ToUpper(strconv.FormatInt(val, 16))
		fmt.Printf("BINARY = %v\n", binary)
		fmt.Printf("HEX = %v\n", hex)

	case "quit", "QUIT":
		fmt.Println(" ")
		fmt.Println(color("GOODBYE👋!!", "\033[1;32m"))
		return

	default:
		fmt.Println(color("NOT IN THE AVAILABLE CONVERSIONS", "\033[1;31m"))
		fmt.Println(color("TRY AGAIN", "\033[1;31m"))
		fmt.Println(" ")
		goto operate
	}
	fmt.Println(" ")
	goto start
}
