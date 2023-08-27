package cmd

import (
	"fmt"
	"io"
	"log"
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
		var input string

		// read text from stdin if flag is specified
		if stdin, _ := cmd.Flags().GetBool("stdin"); stdin {
			buf, err := io.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
			input = string(buf)
		}

		// read text from file if flag is specified
		if fileName, _ := cmd.Flags().GetString("file"); fileName != "" {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed reading file: %s\n", err)
				return
			}
			defer file.Close()

			data, err := io.ReadAll(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed reading file: %s\n", err)
				return
			}

			input = string(data)
		}

		// read text from arguments
		if input == "" && len(args) > 0 {
			input = strings.Join(args, " ")
		}

		// error if no input was supplied
		if input == "" {
			fmt.Fprintln(os.Stderr, "No input was specified")
			return
		}

		output := util.ConvertText(input, false)

		if fileName, _ := cmd.Flags().GetString("output"); fileName != "" {
			data := []byte(output)
			err := os.WriteFile(fileName, data, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed writing file: %s\n", err)
				return
			}
		} else {
			fmt.Println(output)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	decodeCmd.Flags().BoolP("stdin", "s", false, "Read input from standard input")
	decodeCmd.Flags().StringP("file", "f", "", "Read input from file")
	decodeCmd.MarkFlagsMutuallyExclusive("stdin", "file")

	decodeCmd.Flags().StringP("output", "o", "", "Write output to file")
}
