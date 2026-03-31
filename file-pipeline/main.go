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

func trimSpace(s string) string{
	space1 := regexp.MustCompile(`^\s+`)
	space2 := regexp.MustCompile(`$\s+`)

	s = space1.ReplaceAllString(s, "")
	s =space2.ReplaceAllString(s, "")

	return s
}

func replace(s string) string{
	s = strings.ReplaceAll(s, "TODO", "ACTION")
	return s
}

func toTitle(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		for _, char := range words {
			if char >= "A" && char <= "Z" {
				words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
			}
		}
	}
	return strings.Join(words, " ")
}

func toLower(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		for _, char := range words {
			if char >= "a" && char <= "z" {
				words[i] = strings.ToUpper(words[i])
			}
		}
	}
	return strings.Join(words, " ")
}

func flag(s string) string {
	if len(s) > 80 {
		return s + "TRUNCATED"
	}
	return s
}

func dashesAndBlankSpaces(s string) string{
	space := regexp.MustCompile(`\s+`)
	dashes := regexp.MustCompile(`-|_`)

	s = space.ReplaceAllString(s, "")
	s = dashes.ReplaceAllString(s, "")

	return s
}
func main() {
	input, err := os.Open("input.txt")

	if err != nil{
		fmt.Println("FILE NOT FOUND: input.txt")
		return
	}
	
	defer input.Close()

	scanner := bufio.NewScanner(input)
	
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	

}
