package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rokuosan/qrg/internal/clipboard"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

type CommandParameters struct {
	format    string
	level     string
	output    string
	clipboard bool
	size      int
	version   bool
}

var params CommandParameters

var (
	version = "dev"
	commit  = "none"
)

func init() {
	rootCmd.Flags().StringVarP(&params.output, "output", "o", "", "Output file name")
	rootCmd.Flags().StringVarP(&params.format, "format", "", "20060102_15-04-05", "format of the output file")
	rootCmd.Flags().StringVarP(&params.level, "level", "l", "M", "Error Recovery level (L, M, Q, H or 7, 15, 25, 30)")
	rootCmd.Flags().BoolVarP(&params.clipboard, "clipboard", "c", false, "Copy to clipboard")
	rootCmd.Flags().IntVarP(&params.size, "size", "s", 256, "QR code size")
	rootCmd.Flags().BoolVar(&params.version, "version", false, "Show version information")

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

		recoveryLevel, err := parseRecoveryLevel(params.level)
		if err != nil {
			fmt.Println(err)
			return
		}

		qr, err := qrcode.New(args[0], recoveryLevel)
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

// parseRecoveryLevel は、文字列からqrcode.RecoveryLevelを解析する
// 有効なlevelは、L, M, Q, H(大文字・小文字区別なし) または 7, 15, 25, 30
func parseRecoveryLevel(level string) (qrcode.RecoveryLevel, error) {
	level = strings.ToUpper(level)
	switch level {
	case "L", "LOW", strconv.Itoa(int(qrcode.Low)), "7":
		return qrcode.Low, nil
	case "M", "MEDIUM", strconv.Itoa(int(qrcode.Medium)), "15":
		return qrcode.Medium, nil
	case "Q", "QUARTILE", strconv.Itoa(int(qrcode.High)), "25":
		return qrcode.High, nil
	case "H", "HIGHEST", strconv.Itoa(int(qrcode.Highest)), "30":
		return qrcode.Highest, nil
	default:
		return qrcode.Medium, fmt.Errorf("invalid error correction level: %s", level)
	}
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
