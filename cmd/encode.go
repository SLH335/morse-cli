package cmd

import (
	"fmt"
	"os"
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
		output := util.ConvertText(input, true)

		fmt.Println(output)
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
