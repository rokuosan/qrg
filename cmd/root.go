package cmd

import (
	"os"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

var format string
var outFileName string

var rootCmd = &cobra.Command{
	Use:   "qrg",
	Short: "QR Code Generator",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		now := time.Now()

		if outFileName == "" {
			// 出力ファイル名が指定されていない場合は、ファイル名を作ってやる
			outFileName = now.Format(format) + ".png"
		}

		text := args[0]
		if err := qrcode.WriteFile(text, qrcode.Medium, 256, outFileName); err != nil {
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
	rootCmd.Flags().StringVarP(&outFileName, "output", "o", "", "Output file name")
	rootCmd.Flags().StringVarP(&format, "format", "", "20060102_15-04-05", "format of the output file")
}
