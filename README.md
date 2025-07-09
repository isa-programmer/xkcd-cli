# xkcd-cli

[![Go](https://img.shields.io/badge/Go-1.22-blue)](https://golang.org)
[![GitHub release](https://img.shields.io/github/v/release/isa-programmer/xkcd-cli)](https://github.com/isa-programmer/xkcd-cli/releases)
[![License: GPL](https://img.shields.io/badge/License-GPL-blue.svg)](LICENSE)

A command-line interface for [xkcd.com](https://xkcd.com/) comics.

## Features

- Fetch and display XKCD comics right in your terminal
- Cross-platform support: Linux, macOS (Intel and Apple Silicon)
- Simple and fast


## Example screenshot
![Screenshot of xkcd-cli in action](assets/example.webp)

## Installation

### Download pre-built binaries

Grab the latest release for your platform from the [Releases page](https://github.com/isa-programmer/xkcd-cli/releases):
- **Linux (amd64)**
- **macOS (amd64)**
- **macOS (ARM64)**

Example for Linux:

```sh
sudo wget -O /usr/local/bin/xkcd https://github.com/isa-programmer/xkcd-cli/releases/latest/download/xkcd-linux-amd64
sudo chmod +x /usr/local/bin/xkcd
```

Examples for Mac(Apple Silicon/Intel)

```sh
# For Apple Silicon(M1/M2)
sudo wget -O /usr/local/bin/xkcd https://github.com/isa-programmer/xkcd-cli/releases/latest/download/xkcd-darwin-arm64
sudo chmod +x /usr/local/bin/xkcd

# For Apple Intel
sudo wget -O /usr/local/bin/xkcd https://github.com/isa-programmer/xkcd-cli/releases/latest/download/xkcd-darwin-amd64
sudo chmod +x /usr/local/bin/xkcd
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

- Show latest comic:

  ```sh
  xkcd
  ```

- Show a random comic:

  ```sh
  xkcd random
  ```

- Fetch a specific comic by number:

  ```sh
  xkcd 927
  ```

## License

This program licensed with GPLv3. See [LICENSE](LICENSE) for details.

## Credits

- [xkcd.com](https://xkcd.com/) for the awesome comics!
