package util

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const plainCharSeparator = ""
const morseCharSeparator = " "
const plainWordSeparator = " "
const morseWordSeparator = " / "
const invalidPlaceholder = "#"

var plainCharset = [...]string{
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

var morseCharset = [...]string{
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

func ConvertText(input string, isEncoding bool) string {
	var inputCharset [len(plainCharset)]string
	var outputCharset [len(plainCharset)]string
	var inputCharSeparator string
	var outputCharSeparator string
	var inputWordSeparator string
	var outputWordSeparator string

	var invalidChars []string

	if isEncoding {
		inputCharset = plainCharset
		outputCharset = morseCharset
		inputCharSeparator = plainCharSeparator
		outputCharSeparator = morseCharSeparator
		inputWordSeparator = plainWordSeparator
		outputWordSeparator = morseWordSeparator

		input = FormatPlainText(input)
	} else {
		inputCharset = morseCharset
		outputCharset = plainCharset
		inputCharSeparator = morseCharSeparator
		outputCharSeparator = plainCharSeparator
		inputWordSeparator = morseWordSeparator
		outputWordSeparator = plainWordSeparator

		input = FormatMorseCode(input)
	}

	inputLines := strings.Split(input, "\n")

	var outputLines []string
	for _, inputLine := range inputLines {

		inputWords := strings.Split(inputLine, inputWordSeparator)

		var outputWords []string
		for _, inputWord := range inputWords {
			inputWord = strings.TrimSpace(inputWord)
			inputChars := strings.Split(inputWord, inputCharSeparator)

			var outputChars []string
			for _, inputChar := range inputChars {
				inputChar = strings.TrimSpace(inputChar)
				outputChar := invalidPlaceholder
				for i, c := range inputCharset {
					if c == inputChar {
						outputChar = outputCharset[i]
					}
				}

				// check for invalid characters and add them to slice
				if outputChar == invalidPlaceholder && !slices.Contains(invalidChars, inputChar) {
					invalidChars = append(invalidChars, inputChar)
				}

				outputChars = append(outputChars, outputChar)
			}

			outputWord := strings.Join(outputChars, outputCharSeparator)
			outputWords = append(outputWords, outputWord)
		}

		outputLine := strings.Join(outputWords, outputWordSeparator)
		outputLines = append(outputLines, outputLine)
	}

	outputText := strings.Join(outputLines, "\n")

	// if invalid characters are found, display message
	if len(invalidChars) > 0 {
		fmt.Fprintf(os.Stderr, "Invalid characters were found: %s; displaying as \"%s\"\n\n", strings.Join(invalidChars, ", "), invalidPlaceholder)
	}

	return outputText
}
