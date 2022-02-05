# FCalc
> Calculator for fractional numbers

Command line calculator for fractions. The calculator can handle pure fractions, mixed fractions
and negative values. It has zero external dependencies other than `github.com/stretchr/testify` for unit tests.

### Requirements

- Terminal
- Go v1.14+

## Usage

If you have built the project from source or are using the prebuilt binaries simply
run the application from the terminal.

```
./build/fcalc
> Enter a fraction problem:
> ? 1/2 * 3_3/4
> = 1_7/8 
```

## Building the Project

**Compile for Windows x64**

```bash
GOOS=windows GOARCH=amd64 go build -o build/fcalc.exe main.go
```

**Compile for macOS x64**

```bash
GOOS=darwin GOARCH=amd64 go build -o build/fcalc main.go
```

**Compile for Linux x64**

```bash
GOOS=linux GOARCH=amd64 go build -o build/fcalc main.go
```

## Unit Tests

You can run all the projects unit/e2e tests with the following command.

```bash
go test ./...
```