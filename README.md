# qrg

A Simple QRCode Generator

## Installation

```bash
go install github.com/rokuosan/qrg@latest
```

## Usage

```bash
qrg <text>
```

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

## License

MIT
