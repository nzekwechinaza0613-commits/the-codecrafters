package command

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func baseConverter() {

	reader := bufio.NewReader(os.Stdin)

start:
	fmt.Print("C&C>")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.ToLower(input) == "quit" {
		fmt.Print(" Shutting down String Transformer.\n Goodbye.")
		return
	}

	parts := strings.Fields(input)
	convertor := parts[0]

	textSlice := parts[1:]

	text := strings.Join(textSlice, " ")

	if len(text) <= 2 {
		fmt.Println(" ")
		fmt.Println("USAGE: <command> <a> <b>")
		fmt.Println(" ")
		goto start
	}

	switch strings.ToLower(convertor) {

	case "hex":
		val, err := strconv.ParseInt(text, 16, 64)

		if err != nil {
			fmt.Println("VALUE IS NOT A HEXADECIMAL NUMBER")
			fmt.Println("TRY AGAIN")
			fmt.Println(" ")
			goto start
		}
		hex := strings.ToUpper(strconv.FormatInt(val, 10))
		fmt.Printf("DECIMAL = %v\n", hex)

	case "bin":
		val, err := strconv.ParseInt(text, 2, 64)

		if err != nil {
			fmt.Println("VALUE IS NOT A BINARY NUMBER")
			fmt.Println("TRY AGAIN")
			fmt.Println(" ")
			goto start
		}
		fmt.Printf("DECIMAL = %v\n", val)

	case "dec":
		val, err := strconv.ParseInt(text, 10, 64)

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

	default:
		fmt.Println("NOT IN THE AVAILABLE CONVERSIONS")
		fmt.Println("TRY AGAIN")
		fmt.Println(" ")
		goto start
	}
}
