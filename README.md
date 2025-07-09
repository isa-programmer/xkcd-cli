# xkcd-cli

[![Go](https://img.shields.io/badge/Go-1.22-blue)](https://golang.org)
[![GitHub release](https://img.shields.io/github/v/release/isa-programmer/xkcd-cli)](https://github.com/isa-programmer/xkcd-cli/releases)
[![License: GPL](https://img.shields.io/badge/License-GPL-blue.svg)](LICENSE)

A command-line interface for [xkcd.com](https://xkcd.com/) comics.

## Features

- Fetch and display XKCD comics right in your terminal
- Cross-platform support: Linux, macOS (Intel and Apple Silicon)
- Simple and fast

## Installation
>You need to install [feh](https://github.com/derf/feh) for Image viewing
### Download pre-built binaries

Grab the latest release for your platform from the [Releases page](https://github.com/isa-programmer/xkcd-cli/releases):
- **Linux (amd64)**
- **macOS (amd64)**
- **macOS (ARM64)**

Example for Linux:

```sh
wget https://github.com/isa-programmer/xkcd-cli/releases/latest/download/xkcd-linux-amd64
chmod +x xkcd-linux-amd64
./xkcd-linux-amd64
```

### Build from source

Requires Go 1.22+:

```sh
cd xkcd
go build -o xkcd
./xkcd
```

## Usage

```sh
xkcd [options]
```

### Example commands

- Show a random comic:

  ```sh
  xkcd
  ```

- Fetch a specific comic by number:

  ```sh
  xkcd 927
  ```

## Development

The main code lives in the `xkcd/` directory.

```sh
tree
.
├── LICENSE
└── xkcd
    ├── go.mod
    ├── go.sum
    └── main.go
```

You can run tests with:

```sh
cd xkcd
go test ./...
```

## License

This program licensed with GPLv3. See [LICENSE](LICENSE) for details.

## Credits

- [xkcd.com](https://xkcd.com/) for the awesome comics!
