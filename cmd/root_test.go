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

func Test_getRecoveryLevel(t *testing.T) {
	tests := []struct {
		name     string
		recovery string
		want     qrcode.RecoveryLevel
	}{
		{
			name:     "Test Low recovery level with Uppercase",
			recovery: "Low",
			want:     qrcode.Low,
		},
		{
			name:     "Test Low recovery level with large letter",
			recovery: "L",
			want:     qrcode.Low,
		},
		{
			name:     "Test Low recovery level with Lowercase",
			recovery: "low",
			want:     qrcode.Low,
		},
		{
			name:     "Test Medium recovery level with small letter",
			recovery: "l",
			want:     qrcode.Low,
		},
		{
			name:     "Test Low recovery level with Uppercase",
			recovery: "Medium",
			want:     qrcode.Medium,
		},
		{
			name:     "Test Low recovery level with large letter",
			recovery: "M",
			want:     qrcode.Medium,
		},
		{
			name:     "Test Low recovery level with Lowercase",
			recovery: "medium",
			want:     qrcode.Medium,
		},
		{
			name:     "Test Low recovery level with small letter",
			recovery: "m",
			want:     qrcode.Medium,
		},
		{
			name:     "Test High recovery level with Uppercase",
			recovery: "High",
			want:     qrcode.High,
		},
		{
			name:     "Test High recovery level with large letter",
			recovery: "H",
			want:     qrcode.High,
		},
				{
			name:     "Test High recovery level with Lowercase",
			recovery: "high",
			want:     qrcode.High,
		},
				{
			name:     "Test High recovery level with small letter",
			recovery: "h",
			want:     qrcode.High,
		},
		{
			name:     "Test Highest recovery level with Uppercase",
			recovery: "Highest",
			want:     qrcode.Highest,
		},
		{
			name:     "Test Highest recovery level with large letter",
			recovery: "HH",
			want:     qrcode.Highest,
		},
				{
			name:     "Test Highest recovery level with Lowercase",
			recovery: "highest",
			want:     qrcode.Highest,
		},
				{
			name:     "Test Highest recovery level with small letter",
			recovery: "hh",
			want:     qrcode.Highest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRecoveryLevel(tt.recovery); got != tt.want {
				t.Errorf("getRecoveryLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
