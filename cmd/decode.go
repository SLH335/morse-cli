package cmd

import (
	"fmt"
	"os"
	"strings"

	"morse/util"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"dec"},
	Short:   "Convert morse code to plain text",
	Long:    `Convert morse code to plain text`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "Provide at least one argument")
			return
		}

		input := strings.Join(args, " ")
		plainText := decodeMorse(input)

		fmt.Println(plainText)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func decodeMorse(morseText string) string {
	morseText = util.FormatMorseCode(morseText)
	morseWords := strings.Split(morseText, " / ")
	fmt.Println(morseText)

	var plainWords []string
	for _, morseWord := range morseWords {
		morseWord = strings.TrimSpace(morseWord)
		morseLetters := strings.Split(morseWord, " ")

		var plainLetters []string
		for _, morseLetter := range morseLetters {
			morseLetter = strings.TrimSpace(morseLetter)
			plainLetter := decodeMorseLetter(morseLetter)

			plainLetters = append(plainLetters, plainLetter)
		}

		plainWord := strings.Join(plainLetters, "")
		plainWords = append(plainWords, plainWord)
	}

	plainText := strings.Join(plainWords, " ")

	return plainText
}

func decodeMorseLetter(morseLetter string) string {

	for i, l := range util.MorseLetters {
		if l == morseLetter {
			return util.PlainLetters[i]
		}
	}

	return "#"
}
