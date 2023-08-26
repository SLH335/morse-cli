package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"morse/util"

	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:     "encode",
	Aliases: []string{"enc"},
	Short:   "Convert plain text to morse code",
	Long:    `Convert plain text to morse code`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Provide at least one argument")
			return
		}

		input := strings.Join(args, " ")
		morseCode := encodeMorse(input)

		fmt.Println(morseCode)
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func encodeMorse(plainText string) string {
	plainText = util.FormatPlainText(plainText)
	plainWords := strings.Split(plainText, " ")

	var invalidLetters []string

	var morseWords []string
	for _, plainWord := range plainWords {
		plainWord = strings.TrimSpace(plainWord)
		plainLetters := strings.Split(plainWord, "")

		var morseLetters []string
		for _, plainLetter := range plainLetters {
			morseLetter := encodeMorseLetter(plainLetter)

			// check for invalid letters and add them to slice
			if morseLetter == "#" && !slices.Contains(invalidLetters, plainLetter) {
				invalidLetters = append(invalidLetters, plainLetter)
			}

			morseLetters = append(morseLetters, morseLetter)
		}

		morseWord := strings.Join(morseLetters, " ")
		morseWords = append(morseWords, morseWord)
	}

	morseText := strings.Join(morseWords, " / ")

	// if invalid letters are found, display message
	if len(invalidLetters) > 0 {
		fmt.Fprintf(os.Stderr, "Invalid characters were found: %s; displaying as \"#\"\n\n", strings.Join(invalidLetters, ", "))
	}

	return morseText
}

func encodeMorseLetter(plainLetter string) string {

	for i, l := range util.PlainLetters {
		if l == plainLetter {
			return util.MorseLetters[i]
		}
	}

	return "#"
}
