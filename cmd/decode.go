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
				fmt.Printf("Failed reading file: %s\n", err)
				return
			}
			defer file.Close()

			data, err := io.ReadAll(file)
			if err != nil {
				fmt.Printf("Failed reading file: %s\n", err)
				return
			}

			input = string(data)
		}

		// read text from arguments
		if input == "" && len(args) > 0 {
			input = strings.Join(args, " ")
		}

		if input == "" {
			fmt.Fprintln(os.Stderr, "No input was specified")
			return
		}

		output := util.ConvertText(input, false)

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	decodeCmd.Flags().BoolP("stdin", "s", false, "Read text from standard input")
	decodeCmd.Flags().StringP("file", "f", "", "Read text from file")
	decodeCmd.MarkFlagsMutuallyExclusive("stdin", "file")
}
