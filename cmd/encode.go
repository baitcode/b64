/*
Copyright © 2023 Ilia Batii
*/
package cmd

import (
	"encoding/base64"

	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var (
	encodeCmd = &cobra.Command{
		Use:   "encode [data to encode]",
		Short: "Encode data to base64 and print it out",
		Long:  `Encode data to base64 and print it out`,
		Args:  validateInput,
		Run: func(cmd *cobra.Command, args []string) {
			encoded := base64.StdEncoding.EncodeToString([]byte(data))
			cmd.OutOrStdout().Write([]byte(encoded + "\n"))
		},
	}
)

func init() {
	rootCmd.AddCommand(encodeCmd)
}
