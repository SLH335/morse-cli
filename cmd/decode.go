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
		output := util.ConvertText(input, false)

		fmt.Println(output)
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
