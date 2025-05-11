package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

type CommandParameters struct {
	format    string
	output    string
	clipboard bool
}

var params CommandParameters

func init() {
	rootCmd.Flags().StringVarP(&params.output, "output", "o", "", "Output file name")
	rootCmd.Flags().StringVarP(&params.format, "format", "", "20060102_15-04-05", "format of the output file")
	rootCmd.Flags().BoolVarP(&params.clipboard, "clipboard", "c", false, "Copy to clipboard")
}

var rootCmd = &cobra.Command{
	Use:   "qrg",
	Short: "QR Code Generator",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		if err := clipboard.Init(); err != nil {
			panic(err)
		}

		data, err := qrcode.Encode(args[0], qrcode.Medium, 256)
		if err != nil {
			fmt.Println("Failed to generate PNG:", err)
			return
		}

		if params.clipboard {
			clipboard.Write(clipboard.FmtImage, data)
			fmt.Println("Copied to clipboard")
			return
		}

		outFileName := createFileName(params.output, params.format, "png")
		file, err := os.Create(outFileName)
		if err != nil {
			fmt.Println("Failed to create file:", err)
			return
		}
		defer file.Close()
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
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
