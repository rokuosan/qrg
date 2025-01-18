package cmd

import (
	"os"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

var format string

var rootCmd = &cobra.Command{
	Use:   "qrg",
	Short: "QR Code Generator",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		now := time.Now()
		formatted := now.Format(format)

		text := args[0]
		if err := qrcode.WriteFile(text, qrcode.Medium, 256, formatted+".png"); err != nil {
			panic(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&format, "format", "f", "20060102_15-04-05", "format of the output file")
}
