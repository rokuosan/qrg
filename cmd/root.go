package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

type CommandParameters struct {
	format string
	output string
}

var params CommandParameters

// createFileName は、ファイル名を作成する
// extension は、ピリオドなしの拡張子を指定する
func createFileName(fileName string, format string, extension string) string {

	outFileName := fileName

	now := time.Now()
	if fileName == "" {
		// 出力ファイル名が指定されていない場合は、ファイル名を作ってやる
		outFileName = fmt.Sprintf("%s.%s", now.Format(format), extension)
	}

	return outFileName
}

var rootCmd = &cobra.Command{
	Use:   "qrg",
	Short: "QR Code Generator",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}

		outFileName := createFileName(params.output, params.format, "png")

		text := args[0]
		if err := qrcode.WriteFile(text, qrcode.Medium, 256, outFileName); err != nil {
			panic(err)
		}

		fmt.Println(outFileName)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&params.output, "output", "o", "", "Output file name")
	rootCmd.Flags().StringVarP(&params.format, "format", "", "20060102_15-04-05", "format of the output file")
}
