/*
Copyright © 2023 Ilia Batii
*/
package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	data string

	validateInput = func(cmd *cobra.Command, args []string) error {
		var inputReader io.Reader = cmd.InOrStdin()

		fi, _ := os.Stdin.Stat()

		if (fi.Mode() & os.ModeNamedPipe) != 0 {
			inputBytes, err := io.ReadAll(inputReader)
			if err == nil && len(inputBytes) > 0 {
				// pipe input has a trailing newline

				data = string(removeTrainlingNewline(inputBytes))
				args = append(args, data)
			}
		}

		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		data = args[0]
		return nil
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "b64",
	Short: "Cli tool to perform b64 encryption and decryption",
	Long: `Cli tool written with Golang to perform b64 encryption and decryption. 
	Nothing special, just a tool that is supposed to work the same on Linux, Max and Windows 
	for the sake of writing simpler instructions for the software I make. 
	Because, you know. You can't just use base64 encoding in shell in windows.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
