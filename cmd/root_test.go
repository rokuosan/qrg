package cmd

import (
	"testing"
	"time"

	"github.com/skip2/go-qrcode"
)

func Test_createFileName(t *testing.T) {
	type args struct {
		now     time.Time
		outFile string
		format  string
		ext     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with default values",
			args: args{
				now:     time.Date(2025, 5, 11, 0, 0, 0, 0, time.UTC),
				outFile: "",
				format:  "20060102_15-04-05",
				ext:     "png",
			},
			want: "20250511_00-00-00.png",
		},
		{
			name: "Test with custom output file name",
			args: args{
				now:     time.Date(2025, 5, 11, 0, 0, 0, 0, time.UTC),
				outFile: "custom_name",
				format:  "20060102_15-04-05",
				ext:     "png",
			},
			want: "custom_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFileName(tt.args.now, tt.args.outFile, tt.args.format, tt.args.ext); got != tt.want {
				t.Errorf("createFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWriteCloser(t *testing.T) {
	type args = getWriterCloserInput
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test with clipboard",
			args: args{
				isClipboard: true,
				filename:    "",
			},
		},
		{
			name: "Test with file",
			args: args{
				isClipboard: false,
				filename:    "test.png",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getWriteCloser(tt.args)
			if err != nil {
				t.Errorf("getWriteCloser() error = %v", err)
				return
			}
			if got == nil {
				t.Errorf("getWriteCloser() = nil")
			}
		})
	}
}

func Test_parseRecoveryLevel(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantLevel qrcode.RecoveryLevel
		wantErr   bool
	}{
		{"L uppercase", "L", qrcode.Low, false},
		{"l lowercase", "l", qrcode.Low, false},
		{"LOW", "Low", qrcode.Low, false},
		{"0", "0", qrcode.Low, false},

		{"M uppercase", "M", qrcode.Medium, false},
		{"m lowercase", "m", qrcode.Medium, false},
		{"MEDIUM", "Medium", qrcode.Medium, false},
		{"1", "1", qrcode.Medium, false},

		{"Q uppercase", "Q", qrcode.High, false},
		{"q lowercase", "q", qrcode.High, false},
		{"QUARTILE", "Quartile", qrcode.High, false},
		{"2", "2", qrcode.High, false},

		{"H uppercase", "H", qrcode.Highest, false},
		{"h lowercase", "h", qrcode.Highest, false},
		{"HIGHEST", "Highest", qrcode.Highest, false},
		{"3", "3", qrcode.Highest, false},

		{"invalid", "X", qrcode.Medium, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseRecoveryLevel(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseRecoveryLevel(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if got != tt.wantLevel {
				t.Errorf("parseRecoveryLevel(%q) = %v, want %v", tt.input, got, tt.wantLevel)
			}
		})
	}
}
