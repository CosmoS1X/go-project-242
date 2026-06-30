# Pathsize

[![Go](https://github.com/CosmoS1X/go-project-242/actions/workflows/go.yml/badge.svg)](https://github.com/CosmoS1X/go-project-242/actions/workflows/go.yml)

## Overview

Pathsize is a small CLI utility that prints the size of a file or directory.

## Features

- Show the size of a single file or directory
- Support recursive directory traversal
- Optionally include hidden files and directories
- Output sizes in either raw bytes or human-readable format

## Installation

Install the CLI binary with:

```bash
go install ./cmd/pathsize
```

Or download the prebuilt binaries from the [releases page](https://github.com/CosmoS1X/pathsize/releases)

After installation, run:

```bash
pathsize <path>
```

## Library usage

The package can also be used as a Go library by importing it from the module root:

```go
import "github.com/CosmoS1X/pathsize"

size, err := pathsize.Get(path, recursive, human, all)
```

`Get` returns the formatted size as a string and an error if the path cannot be read.

Parameters:

- `path` — the file or directory to inspect
- `recursive` — if `true`, directories are traversed recursively
- `human` — if `true`, the result is formatted in human-readable units
- `all` — if `true`, hidden files and directories are included in the calculation

## Usage

```bash
pathsize [options] <path>
```

### Options

- `-r`, `--recursive` — calculate the size of directories recursively
- `-H`, `--human` — display sizes in a human-readable format
- `-a`, `--all` — include hidden files and directories
- `-h`, `--help` — show help information

## Examples

Show the size of a file in bytes:

```bash
pathsize testdata/test.txt
```

Show a directory size recursively:

```bash
pathsize -r testdata
```

Show a human-readable size for a directory:

```bash
pathsize -H testdata
```

Include hidden entries in the calculation:

```bash
pathsize -a testdata
```

Combine multiple flags in one command:

```bash
pathsize -raH testdata
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
