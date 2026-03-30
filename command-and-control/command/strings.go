package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Upper(s string) string {
	w := strings.Fields(s)

	for i := 0; i < len(w); i++ {

		w[i] = strings.ToUpper(w[i])
	}
	return strings.Join(w, " ")
}

func lower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s)
}

func capitalize(s string) string {
	words := strings.Fields(s)

	for i, item := range words {
		words[i] = strings.ToUpper(string(item[0])) + strings.ToLower(item[1:])

	}
	return strings.Join(words, " ")
}

func title(s string) string {
	words := strings.Fields(s)

	smallWords := map[string]bool{
		"a": true, "an": true, "of": true, "the": true, "and": true, "but": true,
		"or": true, "for": true, "nor": true, "on": true, "at": true, "to": true,
		"by": true, "in": true, "it": true, "is": true, "as": true, "up": true,
	}

	for i, item := range words {
		lower := strings.ToLower(item)

		if i == 0 || !smallWords[lower] {
			words[i] = strings.ToUpper(string(lower[0])) + lower[1:]
		} else {
			words[i] = lower
		}
	}
	return strings.Join(words, " ")
}

func snake(s string) string {

	ToLower := strings.ToLower(s)

	words := strings.Fields(ToLower)

	for i, item := range words {
		updated := ""
		for _, char := range item {
			if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
				updated += string(char)
			}
		}
		words[i] = updated
	}
	return strings.Join(words, "_")
}

func reverse(s string) string {
	rev := []rune(s)
	var reverse []rune
	for i := len(rev) - 1; i >= 0; i-- {
		reverse = append(reverse, rev[i])
	}
	return string(reverse)
}

func strTransformer() {
start:
	fmt.Print("C&C> ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err != nil {
		fmt.Println("error")
		fmt.Println(" ")
	}

	if input == "" {
		goto start
	}
	if strings.ToLower(input) == "quit" {
		fmt.Print(" Shutting down String Transformer.\n Goodbye.")
		return
	}

	parts := strings.Fields(input)

	cases := parts[0]
	textSclice := parts[1:]

	text := strings.Join(textSclice, " ")

	if len(text) <= 1 {
		fmt.Println(" ")
		fmt.Println("NO TEXT PROVIDED. USAGE: <command> <text>")
		fmt.Println(" ")
		goto start
	}

	switch strings.ToLower(cases) {

	case "upper":
		fmt.Println(" ")
		fmt.Println(Upper(text))

	case "lower":
		fmt.Println(" ")
		fmt.Println(lower(text))

	case "cap":
		fmt.Println(" ")
		fmt.Println(capitalize(text))

	case "title":
		fmt.Println(" ")
		fmt.Println(title(text))

	case "snake":
		fmt.Println(" ")
		fmt.Println(snake(text))

	case "reverse":
		fmt.Println(" ")
		fmt.Println(reverse(text))

	default:
		fmt.Println("UNKNOWN COMMAND")
		fmt.Println("TRY AGAIN")
		fmt.Println(" ")

	}

}
