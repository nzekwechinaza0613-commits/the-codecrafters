// package command

// import (
// 	"strconv"
// 	"fmt"
// 	"strings"
// )

// func add(a, b float64) float64 {
// 	return a + b
// }

// func sub(a, b float64) float64 {
// 	return a - b
// }

// func mul(a, b float64) float64 {
// 	return a * b
// }

// func mod(a, b int) int {
// 	return a % b
// }
// func divide(a, b float64) float64 {
// 	return a / b
// }
// func power(a, b int) int {
// 	return a ^ b
// }

// func dectobin(word string) string {
// 	words, _ := strconv.ParseInt(word, 10, 64)
// 	return strconv.FormatInt(words, 2)
// }
// func dectohex(word string) string {
// 	words, _ := strconv.ParseInt(word, 10, 64)
// 	return strconv.FormatInt(words, 16)
// }

// func bintodec(word string) (int64, error) {
// 	words, err := strconv.ParseInt(word, 2, 64)
// 	if err != nil {
// 		fmt.Println("not a valid binary", err)
// 		return words, err
// 	}
// 	return words, err

// }
// func hextodec(word string) (int64, error) {
// 	words, err := strconv.ParseInt(word, 16, 64)
// 	if err != nil {
// 		fmt.Println("not a valid hexadecimal", err)
// 		return words, err
// 	}
// 	return words, err

// }
// func bintohex(word string) (int64, error) {
// 	words, err := strconv.ParseInt(word, 2, 64)
// 	if err != nil {
// 		fmt.Println("not a valid binary", err)
// 		return words, err
// 	}
// 	return words, err

// }
// func upper(word string) string {
// 	return strings.ToUpper(word)
// }
// func lower(word string) string {
// 	return strings.ToLower(word)
// }
// func cap(word string) string {
// 	words := strings.Fields(word)
// 	for i, r := range words {
// 		words[i] = strings.ToUpper(string(r[0])) + strings.ToLower(string(r[1:]))
// 	}

// 	return strings.Join(words, " ")
// }

// func title(s string) string {
// 	words := strings.Fields(s)

// 	smallWords := map[string]bool{
// 		"a": true, "an": true, "of": true, "the": true, "and": true, "but": true,
// 		"or": true, "for": true, "nor": true, "on": true, "at": true, "to": true,
// 		"by": true, "in": true, "it": true, "is": true, "as": true, "up": true,
// 	}

// 	for i, item := range words {
// 		lower := strings.ToLower(item)

// 		if i == 0 || !smallWords[lower] {
// 			words[i] = strings.ToUpper(string(lower[0])) + lower[1:]
// 		} else {
// 			words[i] = lower
// 		}
// 	}
// 	return strings.Join(words, " ")
// }
// func snake(s string) string {

// 	ToLower := strings.ToLower(s)

// 	words := strings.Fields(ToLower)

// 	for i, item := range words {
// 		updated := ""
// 		for _, char := range item {
// 			if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
// 				updated += string(char)
// 			}
// 		}
// 		words[i] = updated
// 	}
// 	return strings.Join(words, "_")
// }

// func reverse(s string) string {
// 	rev := []rune(s)
// 	var reverse []rune
// 	for i := len(rev) - 1; i >= 0; i-- {
// 		reverse = append(reverse, rev[i])
// 	}
// 	return string(reverse)
// }
package command