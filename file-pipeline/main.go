//   	 ═══════════════════════════════════════════
//      SQUAD PIPELINE CONTRACT
//         Squad: [BenchMarkers]
//      ───────────────────────────────────────────
//      Input line types:
//        [line 1: ALL CAP]
//        [line 2: lower]
//        [line 3: Trimspace]
//        [line 4: TODO ]

//      Transformation rules (in order):
//        1. [Trim all leading and trailing whitespace ]
//        2. [Replace TODO: with ✦ ACTION]
//        3. [Convert ALL CAPS lines to Title Case ]
//        4. [Convert all lowercase lines to uppercase ]
//        5. [Remove lines that are only dashes or blanks 	]

//      Output format:
//        [Header: yes/no — SENTINEL FIELD REPORT — PROCESSED]
//        [Line numbering format- 1]

//      Terminal summary fields:
//         ✦ Lines read    : [number]
//        ✦ Lines written : [number]
//       ✦ Lines removed : [number]
//        ✦ Rules applied :
//        1. [Trim all leading and trailing whitespace ]
//        2. [Replace TODO: with ✦ ACTION]
//        3. [Convert ALL CAPS lines to Title Case ]
//        4. [Convert all lowercase lines to uppercase ]
//        5. [Remove lines that are only dashes or blanks 	]
//      ═══════════════════════════════════════════ -->

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func trimSpace(text string) string {
	space1 := regexp.MustCompile(`^\text+`)
	space2 := regexp.MustCompile(`$\text+`)

	text = space1.ReplaceAllString(text, "")
	text = space2.ReplaceAllString(text, "")

	return text
}

func replace(text string) string {
	text = strings.ReplaceAll(text, "TODO", "ACTION")
	return text
}

func toTitle(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		for _, char := range words {
			if char >= "A" && char <= "Z" {
				words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
			}
		}
	}
	return strings.Join(words, " ")
}

func toLower(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		for _, char := range words {
			if char >= "a" && char <= "z" {
				words[i] = strings.ToUpper(words[i])
			}
		}
	}
	return strings.Join(words, " ")
}

func flag(text string) string {
	if len(text) > 80 {
		return text + "TRUNCATED"
	}
	return text
}

func dashesAndBlankSpaces(text string) string {
	space := regexp.MustCompile(`\text+`)
	dashes := regexp.MustCompile(`-|_`)

	text = space.ReplaceAllString(text, "")
	text = dashes.ReplaceAllString(text, "")

	return text
}
func main() {
	fileText := "input.txt"
	input, err := os.Open(fileText)

	if err != nil {
		fmt.Println("FILE NOT FOUND: input.txt")
		return
	}

	defer input.Close()

	file, err := input.Stat()
	if err != nil {
		fmt.Println(err)
	}

	if file.Size() == 0 {
		fmt.Println("⚠Input file is empty. Nothing to process")
		return
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	text := string(data)
	text = trimSpace(text)
	text = replace(text)
	text = toTitle(text)
	text = toLower(text)
	text = flag(text)
	text = dashesAndBlankSpaces(text)

	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}

}
