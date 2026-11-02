# Go Reloaded

A text auto-correction and formatting tool written in Go. It reads a text file, applies a set of editing rules (case changes, number conversion, punctuation and grammar fixes) and writes the corrected text to a new file.

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![CLI](https://img.shields.io/badge/type-CLI-informational)

## Example

**Input** (`sample.txt`)

```text
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was 1E (hex) and 10 (bin) years ago , I was sitting over there ,and then BAMM !! there is a untold story
```

**Output** (`result.txt`)

```text
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was 30 and 2 years ago, I was sitting over there, and then BAMM!! there is an untold story
```

## Supported Rules

| Marker | Effect | Example |
|---|---|---|
| `(hex)` | Converts the previous hexadecimal number to decimal | `1E (hex)` → `30` |
| `(bin)` | Converts the previous binary number to decimal | `10 (bin)` → `2` |
| `(up)` / `(up, n)` | Uppercases the previous word / previous *n* words | `go (up)` → `GO` |
| `(low)` / `(low, n)` | Lowercases the previous word / previous *n* words | `STOP (low)` → `stop` |
| `(cap)` / `(cap, n)` | Capitalizes the previous word / previous *n* words | `bridge (cap)` → `Bridge` |
| Punctuation | `.` `,` `!` `?` `:` `;` attach to the previous word; groups such as `...` and `!?` stay together | `there ,and` → `there, and` |
| Quotes | Single quotes `' '` wrap the enclosed words without inner spaces | `' awesome '` → `'awesome'` |
| Articles | `a` becomes `an` before a vowel or `h` | `a apple` → `an apple` |

## Getting Started

**Requirements:** Go 1.20 or newer

```bash
git clone https://github.com/ahmed-jawher/go-reloaded.git
cd go-reloaded
go run . sample.txt result.txt
```

The output file must have a `.txt` extension. Invalid markers, such as `(hex)` after a non-hexadecimal word, produce a clear error message.

## Project Structure

```text
go-reloaded/
├── main.go              # CLI: argument validation, file I/O
└── pkg/
    ├── check_all.go     # main pipeline: tokenizes text and applies every rule
    ├── up.go, up_x.go   # (up) and (up, n)
    ├── low.go, low_x.go # (low) and (low, n)
    ├── cap.go, cap_x.go # (cap) and (cap, n)
    ├── punctuation.go   # punctuation and quote formatting
    └── helpers.go       # shared helpers
```

## Author

**Ahmed Jawher** — [@ahmed-jawher](https://github.com/ahmed-jawher)
