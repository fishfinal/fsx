# fsx

<div align="center">

<img src="https://img.shields.io/badge/Go-v1.26.6-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version" />
<img src="https://img.shields.io/badge/license-MIT-green?style=for-the-badge" alt="License" />
<img src="https://img.shields.io/badge/release-v1.0.0-blue?style=for-the-badge" alt="Release" />
<img src="https://img.shields.io/badge/build-passing-brightgreen?style=for-the-badge&logo=githubactions&logoColor=white" alt="Build Status" />
<img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge" alt="PRs Welcome" />

</div>

A Go filesystem extension package that provides convenient helpers for common file and directory operations: absolute path resolution, blank-line cleanup, content trimming by markers, line/character counting, embedded filesystem helpers, ZIP archiving, HTTP file download/upload, and more.

## Installation

```shell
go get -u github.com/fishfinal/fsx
```

## Features

- [x] Absolute path resolution
- [x] Remove blank lines from files or `[]byte` data
- [x] Remove content before/after a specified marker string
- [x] Count lines and characters in files or `[]byte` data
- [x] File and directory existence checks (including `embed.FS`)
- [x] Create, truncate, touch, copy, and read directories
- [x] Filename and extension utilities
- [x] Read file content as `[]byte` or line slices
- [x] Extract lines matching a regular expression
- [x] Remove lines matching a regular expression
- [x] ZIP archive creation from a file or directory
- [x] HTTP file download helpers (with RFC 5987 filename encoding)
- [x] HTTP upload MIME type validation
- [x] Response header helpers (`Content-Disposition` filename parsing)
- [x] Go version parsing (`VersionGithub`)

## Quick Start

### Absolute path

```go
path, err := fsx.AbsPath("./data", true) // create the directory if it does not exist
if err != nil {
    log.Fatal(err)
}
fmt.Println(path)
```

### Clean blank lines

```go
// From bytes
cleaned := fsx.CleanBlankLines([]byte("hello\n\nworld\n"))

// From a file
cleaned, err := fsx.CleanBlankLinesForFile("input.txt")
```

### Trim content by marker

```go
data, _ := os.ReadFile("article.txt")

// Keep content from the marker onward (marker included by default)
prefix := fsx.MarkContentPrefixClear(data, "START")

// Keep content up to the marker (marker excluded by default)
suffix := fsx.MarkContentSuffixClear(data, "END")

// Control whether the marker string itself is kept
prefix = fsx.MarkContentPrefixClear(data, "START",
    fsx.WithSpecificContainsMarkString(false))
```

### Count lines and characters

```go
lines := fsx.Lines([]byte("a\nb\nc"))           // 3
chars := fsx.Characters([]byte("a\nb\nc"))      // 3 (blank lines ignored)

lines, err := fsx.LinesForFile("input.txt")
chars, err = fsx.CharactersForFile("input.txt")
```

### File existence

```go
exists := fsx.Exists("config.yaml")

// Works with embed.FS as well
//go:embed static/*
var staticFS embed.FS

exists = fsx.EmbedsFileExists(staticFS, "static/css/style.css")
exists = fsx.EmbedsExists(staticFS, "static/images/")
```

### Create a file

```go
path, err := fsx.CreateFile("report", []byte("hello"),
    fsx.WithSpecificBaseDir("output"),
    fsx.WithSpecificFileSystemPath("docs"),
    fsx.WithSpecificFileExtension(fsx.FileExtensionWithDocumentOfMarkdown),
)
// output/docs/static/report.md
```

### Touch / truncate / copy

```go
_ = fsx.Touch("logs/app.log")            // creates parent dirs if needed
_ = fsx.Truncate("logs/app.log", 0)      // truncate to zero bytes

files, _ := fsx.ReadDir("src", nil)
_ = fsx.WriteDir("dst", files, fsx.WithOriginalFileNameTrimPrefix("src"))
```

### Read file content as line slices

```go
lines, err := fsx.GetFileContentStringSliceWithFilename("input.txt",
    fsx.WithIgnoreBlankLine(true))

lines, err = fsx.GetFileContentStringSliceWithBuffer(buf)

lines, err = fsx.GetFileContentStringSliceWithEmbedFS(staticFS, "static/readme.md")
```

### Regular expression helpers

```go
// Extract matching lines
matches, err := fsx.GetContentWithRegularExpression(content, `\d+`)

// Remove matching lines
cleaned := fsx.CleanSpecificContentLines(data, `^#`)
```

### ZIP archive

```go
// Compress a directory or a single file. The archive is placed next to the source.
zipPath, err := fsx.CreateZIPFileWithFilename("./test")
// -> ./test.zip

zipPath, err = fsx.CreateZIPFileWithFilename("./main.go")
// -> ./main.zip
```

### HTTP download

```go
func handler(w http.ResponseWriter, r *http.Request) {
    if err := fsx.DownloadFile(w, "files/report.pdf"); err != nil {
        log.Println(err)
    }
}

// Or stream from an io.Reader (pass -1 to auto-detect size for io.ReadSeeker)
func streamHandler(w http.ResponseWriter, r *http.Request) {
    file, _ := os.Open("files/report.pdf")
    defer file.Close()

    _ = fsx.DownloadFileByIOReader(w, file, "报告.pdf", -1)
}
```

Filenames are encoded per RFC 5987 (`filename*=UTF-8''...`), so non-ASCII names are transmitted safely.

### HTTP upload MIME check

```go
allowed, err := fsx.UploadIsAllowMIMEWithFormFile(r, "avatar",
    []string{"image/jpeg", "image/png"})
```

### Response header helpers

```go
// Extract the filename from a Content-Disposition header
name := fsx.GetResponseHeaderFunc(w, "Content-Disposition",
    fsx.ContentDispositionFilename())
```

### Go version parsing

```go
v, err := fsx.GetVersion(runtime.Version()) // e.g. "go1.21.0"
fmt.Println(v.GetFullVersion())  // 1.21.0
fmt.Println(v.GetMinorVersion()) // 1.21
```

## API Overview

| File | Description |
| --- | --- |
| `abs.go` | `AbsPath` – resolve and optionally create absolute paths |
| `blank_line.go` | `CleanBlankLines`, `CleanBlankLinesForFile` |
| `content.go` | `MarkContentPrefixClear`, `MarkContentSuffixClear` |
| `create.go` | `CreateFile` |
| `directory.go` | `Exists`, `EmbedsExists`, `MkdirAll`, `ReadDir`, `WriteDir`, `Copy`, `IsDir`, `DirNotExists`, `LastDirName`, … |
| `download.go` | `DownloadFile`, `DownloadFileByIOReader` |
| `extension.go` | `FileExtension` type and predefined constants |
| `file.go` | `CreateFilename`, `GetFileExtension`, `IsFile`, `Touch`, `Truncate`, `GetFileContent*`, `GetContentWithRegularExpression`, … |
| `line.go` | `Lines`, `LinesForFile`, `Characters`, `CharactersForFile` |
| `option.go` | Functional options (`WithSpecific*`, `WithIgnoreBlankLine`, …) |
| `path.go` | `FileSystemPath` type and predefined paths |
| `remove_line.go` | `CleanSpecificContentLines`, `CleanSpecificContentLinesForFile` |
| `response.go` | `GetResponseHeader`, `ContentDispositionFilename` |
| `slice.go` | `InSlice` |
| `upload.go` | `UploadIsAllowMIMEWithFormFile` |
| `version.go` | `VersionGithub`, `GetVersion` |
| `zip.go` | `CreateZIPFileWithFilename` |

## Testing

```shell
go test ./...
```

Test fixtures live under `test/data/`. The suite uses [`github.com/stretchr/testify`](https://github.com/stretchr/testify) for assertions and suite-based organization. Generated artifacts are cleaned up in `TearDownTest`.

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.

---

Copyright 2021 fishfinal &lt;contact.fishfinal@gmail.com&gt;
