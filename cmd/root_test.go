package cmd

import (
	"testing"
	"time"
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
