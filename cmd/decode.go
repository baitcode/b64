/*
Copyright © 2023 Ilia Batii
*/
package cmd

import (
	"os"
	"time"

	"encoding/base64"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode [data to encode]",
	Short: "Decode data to base64 and print it out",
	Long:  `Decode data to base64 and print it out`,
	Args:  validateInput,
	Run: func(cmd *cobra.Command, args []string) {
		decoded, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			cmd.ErrOrStderr().Write([]byte("Error: " + err.Error()))
			return
		}
		decoded = append(decoded, '\n')
		cmd.SetOut(os.Stdout)
		cmd.OutOrStdout().Write(decoded)
		time.Sleep(100 * time.Millisecond)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}
