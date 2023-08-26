package util

import (
	"strings"
)

// format plain text to improve parsing and return
func FormatPlainText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToUpper(text)

	// convert alternative inputs to their standard form
	text = strings.ReplaceAll(text, "×", "x")
	text = strings.ReplaceAll(text, "%", "/")

	return text
}

// format morse code to improve parsing and return
func FormatMorseCode(code string) string {
	code = strings.ReplaceAll(code, "_", "-")

	code = strings.TrimSpace(code)

	code = strings.ReplaceAll(code, "   ", " / ")

	codeWords := strings.Split(code, " / ")
	var formattedWords []string
	for _, codeWord := range codeWords {
		codeWord = strings.TrimSpace(codeWord)

		if codeWord != "" {
			formattedWords = append(formattedWords, strings.TrimSpace(codeWord))
		}
	}
	code = strings.Join(formattedWords, " / ")

	return code
}

var PlainLetters = [...]string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
	"Ä",
	"Ö",
	"Ü",
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"&",
	"'",
	"@",
	")",
	"(",
	":",
	",",
	"=",
	"!",
	".",
	"-",
	"+",
	"\"",
	"?",
	"/",
}

var MorseLetters = [...]string{
	".-",
	"-...",
	"-.-.",
	"-..",
	".",
	"..-.",
	"--.",
	"....",
	"..",
	".---",
	"-.-",
	".-..",
	"--",
	"-.",
	"---",
	".--.",
	"--.-",
	".-.",
	"...",
	"-",
	"..-",
	"...-",
	".--",
	"-..-",
	"-.--",
	"--..",
	".-.-",
	"---.",
	"..--",
	"-----",
	".----",
	"..---",
	"...--",
	"....-",
	".....",
	"-....",
	"--...",
	"---..",
	"----.",
	".-...",
	".----.",
	".--.-.",
	"-.--.-",
	"-.--.",
	"---...",
	"--..--",
	"-...-",
	"-.-.--",
	".-.-.-",
	"-....-",
	".-.-.",
	".-..-.",
	"..--..",
	"-..-.",
}
