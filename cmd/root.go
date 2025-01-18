package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qrg",
	Short: "QR Code Generator",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		now := time.Now()
		formatted := now.Format("20060102_15-04-05")
		filename := fmt.Sprintf("%s.png", formatted)

		text := args[0]
		if err := qrcode.WriteFile(text, qrcode.Medium, 256, filename); err != nil {
			panic(err)
		}

		fmt.Println(filename)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
