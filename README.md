# Path Size

### Hexlet tests and linter status:

[![Actions Status](https://github.com/CosmoS1X/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/CosmoS1X/go-project-242/actions)
[![Go](https://github.com/CosmoS1X/go-project-242/actions/workflows/go.yml/badge.svg)](https://github.com/CosmoS1X/go-project-242/actions/workflows/go.yml)

## Overview

Path Size is a small CLI utility that prints the size of a file or directory.

## Features

- Show the size of a single file or directory
- Support recursive directory traversal
- Optionally include hidden files and directories
- Output sizes in either raw bytes or human-readable format

## Installation

Install the CLI binary with:

```bash
go install ./cmd/hexlet-path-size
```

After installation, run:

```bash
hexlet-path-size <path>
```

## Library usage

The package can also be used as a Go library by importing it from the module root:

```go
import "code"

size, err := code.GetPathSize(path, recursive, human, all)
```

`GetPathSize` returns the formatted size as a string and an error if the path cannot be read.

Parameters:

- `path` — the file or directory to inspect
- `recursive` — if `true`, directories are traversed recursively
- `human` — if `true`, the result is formatted in human-readable units
- `all` — if `true`, hidden files and directories are included in the calculation

## Usage

```bash
hexlet-path-size [options] <path>
```

### Options

- `-r`, `--recursive` — calculate the size of directories recursively
- `-H`, `--human` — display sizes in a human-readable format
- `-a`, `--all` — include hidden files and directories
- `-h`, `--help` — show help information

## Examples

Show the size of a file in bytes:

```bash
hexlet-path-size testdata/test.txt
```

Show a directory size recursively:

```bash
hexlet-path-size -r testdata
```

Show a human-readable size for a directory:

```bash
hexlet-path-size -H testdata
```

Include hidden entries in the calculation:

```bash
hexlet-path-size -a testdata
```

Combine multiple flags in one command:

```bash
hexlet-path-size -raH testdata
```

## Output format

The program prints:

```text
<size>    <path>
```

For example:

```text
6B    testdata/test.txt
```
