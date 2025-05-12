# qrg

A Simple QRCode Generator.

## Installation

```bash
go install github.com/rokuosan/qrg@latest
```

## Usage

```bash
qrg <text>
```

## Options

- `--format` : Specify the datetime format. The default is the Go format `20060102_150405`.
- `-c` `--clipboard` : Save the QR code to the clipboard instead of a file.
- `-o` `--output` : Specify the output file name. (e.g. `qrg -o qrcode.png "text"`)

## Example

When you run the command, a PNG file with the date and time as the file name will be generated in the current directory.

Print the created file name to standard output.

```bash
$ qrg "https://example.com/"
20250118_20-42-26.png
```

On macOS, you can use the open command to view the generated QR code in the Preview app.

```bash
$ qrg "https://example.com/" | xargs open
```

If you want to save the QR code to your clipboard, you can use the `--clipboard`(`-c`) option.

```bash
$ qrg -c "https://example.com/"
```

## License

This software is released under the MIT License, see LICENSE.
