package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"command-and-control/calculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
//start:

	fmt.Println("AVAILABLE COMMANDS: ")
	fmt.Print("1.calculator\n 2.base converter\n 3.string transformer\n")

	fmt.Println("choose a command")
	// input, _ := reader.ReadString('\n')
	// input = strings.TrimSpace(input)
	// if input != "1" && input != "2" && input != "3" {
	// 	fmt.Println("not a command")
	// 	goto start
	// }
	{
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		read := strings.Fields(line)
		if read[0] == "calc"{
			if read[1] != "add" && read[1] != "sub" && read[1] != "mul" && read[1] != "div" && read[1] != "power" && read[1] != "mod" {
				fmt.Println("not a command")
			}
			if read[0] == "add" {
				a, _ := strconv.ParseFloat(read[1], 64)
				b, _ := strconv.ParseFloat(read[1], 64)
				fmt.Println(add(a,b))
			}
			if read[0] == "sub" {
				a, _ := strconv.ParseFloat(read[1], 64)
				b, _ := strconv.ParseFloat(read[1], 64)
				fmt.Println(sub(a,b))
			}
			if read[0] == "mul" {
				a, _ := strconv.ParseFloat(read[1], 64)
				b, _ := strconv.ParseFloat(read[1], 64)
				fmt.Println(mul(a,b))
			}

			if read[0] == "div" {
				a, _ := strconv.ParseFloat(read[1], 64)
				b, _ := strconv.ParseFloat(read[1], 64)
				fmt.Println(divide(a,b))
			}

			if read[0] == "power" {
				a, _ := strconv.Atoi(read[1])
				b, _ := strconv.Atoi(read[1])
				fmt.Println(power(a,b))
			}

			if read[0] == "mod" {
				a, _ := strconv.Atoi(read[1])
				b, _ := strconv.Atoi(read[1])
				fmt.Println(mod(a,b))
			}
		}
	}
	// switch input {
	// case "1":
	// 	fmt.Println(calculator(upper))
	// }

}
