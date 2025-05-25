package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rokuosan/qrg/internal/clipboard"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

type CommandParameters struct {
	format    string
	output    string
	clipboard bool
	size      int
	version   bool
	recovery string
}

var params CommandParameters

var (
	version = "dev"
	commit  = "none"
)

func init() {
	rootCmd.Flags().StringVarP(&params.output, "output", "o", "", "Output file name")
	rootCmd.Flags().StringVarP(&params.format, "format", "", "20060102_15-04-05", "format of the output file")
	rootCmd.Flags().BoolVarP(&params.clipboard, "clipboard", "c", false, "Copy to clipboard")
	rootCmd.Flags().IntVarP(&params.size, "size", "s", 256, "QR code size")
	rootCmd.Flags().BoolVar(&params.version, "version", false, "Show version information")
	rootCmd.Flags().StringVarP(&params.recovery, "recovery", "r", "Medium", "Recovery level (Low, Medium, Quartile, High)")

	params.output = createFileName(time.Now(), params.output, params.format, "png")
}

var rootCmd = &cobra.Command{
	Use:   "qrg",
	Short: "QR Code Generator",
	Run: func(cmd *cobra.Command, args []string) {
		if params.version {
			fmt.Printf("%s (%s)\n", version, commit)
			return
		}
		if len(args) == 0 {
			cmd.Help()
			return
		}
		if err := clipboard.Init(); err != nil {
			panic(err)
		}

		qr, err := qrcode.New(args[0], getRecoveryLevel(params.recovery))
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

		if qr.Write(params.size, w); err != nil {
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

func getRecoveryLevel(level string) qrcode.RecoveryLevel {
	level = strings.ToUpper(level)
	switch level {
	case "LOW", "L":
		return qrcode.Low
	case "MEDIUM", "M":
		return qrcode.Medium
	case "HIGH", "H":
		return qrcode.High
	case "HIGHEST", "HH":
		return qrcode.Highest
	default:
		fmt.Printf("Unknown recovery level: %s, defaulting to Medium\n", level)
		return qrcode.Medium
	}
}
