# 🍕 Go Slices Package Check 🍕

[![Go](https://github.com/manuelarte/goslicespackagecheck/actions/workflows/go.yml/badge.svg)](https://github.com/manuelarte/goslicespackagecheck/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/manuelarte/goslicespackagecheck)](https://goreportcard.com/report/github.com/manuelarte/goslicespackagecheck)
![version](https://img.shields.io/github/v/release/manuelarte/goslicespackagecheck)

This 🧐 linter checks whether some of your functions can be replaced by already existing 🍕 [slices](https://pkg.go.dev/slices) functions.

> [!WARNING]  
> This linter can't guarantee is working for every case, please double check the diagnosis given.

## ⬇️  Getting Started

To install it run:

```bash
go install github.com/manuelarte/goslicespackagecheck@latest
```

To run it:

```bash
goslicespackagecheck [-concat=true|false] [-equal=true|false] [-max=true|false]
```

- `concat`: (default `true`) To enable/disable `slices.Concat` check.
- `equal`: (default `true`) To enable/disable `slices.Equal` check.
- `max`: (default `true`) To enable/disable `slices.Max` check.

## 🚀 Features

### slices.Concat

Detect for loops that can be replaced by [`slices.Concat`](https://pkg.go.dev/maps#Clone). Example can be found [here](examples/slices/concat/main.go).

### slices.Equal

Detect functions that can be replaced by [`slices.Equal`](https://pkg.go.dev/slices#Equal). Example can be found [here](./examples/slices/equal/main.go).

### slices.Max

Detect for loops that can be replaced by [`slices.Max`](https://pkg.go.dev/slices#Max). Example can be found [here](./examples/slices/max/main.go).
