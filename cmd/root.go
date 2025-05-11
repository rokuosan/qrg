package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rokuosan/qrg/internal/clipboard"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
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

	params.output = createFileName(time.Now(), params.output, params.format, "png")
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

		qr, err := qrcode.New(args[0], qrcode.Medium)
		if err != nil {
			fmt.Println("Failed to generate PNG:", err)
			return
		}

		w, err := getWriteCloser(getWriterCloserInput{
			isClipboard: params.clipboard,
			filename:    params.output,
		})
		if err != nil {
			fmt.Println("Failed to create writer:", err)
			return
		}
		defer w.Close()

		if qr.Write(256, w); err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}

		if params.clipboard {
			fmt.Println("Copied to clipboard")
		} else {
			fmt.Println(params.output)
		}
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
func createFileName(now time.Time, fileName string, format string, extension string) string {
	if fileName != "" {
		return fileName
	}
	// 出力ファイル名が指定されていない場合は、ファイル名を作ってやる
	return fmt.Sprintf("%s.%s", now.Format(format), extension)
}

type getWriterCloserInput struct {
	isClipboard bool
	filename    string
}

func getWriteCloser(input getWriterCloserInput) (io.WriteCloser, error) {
	if input.isClipboard {
		return clipboard.New(), nil
	}
	f, err := os.Create(input.filename)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return nil, err
	}
	return f, nil
}
