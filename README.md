# kebab-case

## Description

This repository contains a utility for converting strings to kebab-case format. Kebab-case is a lowercase string format where words are separated by hyphens (`-`). It is commonly used in URLs, file names, and CSS class names.

## Usage

### Command line

```sh
go run . "Hello World"
# hello-world

go run . getHTTPResponseCode
# get-http-response-code
```

### In Go

```go
result := ToKebabCase("Hello World")
fmt.Println(result) // hello-world
```

## API

### `ToKebabCase(s string) string`

Converts the given string to kebab-case.

Words are split on:

- spaces, punctuation, underscores and existing hyphens (`hello_world` → `hello-world`)
- lower-to-upper case changes (`helloWorld` → `hello-world`)
- the end of an acronym (`HTTPServer` → `http-server`)
- changes between letters and digits (`version2` → `version-2`)

Repeated or leading/trailing separators are collapsed, and non-ASCII letters are supported (`ÉcoleNormale` → `école-normale`).

## Running the tests

```sh
go test ./...
```
