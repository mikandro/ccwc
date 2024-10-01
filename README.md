# ccwc

A command-line word counter tool written in Go. `ccwc` is similar to the Unix `wc` command and can count the number of bytes, lines, words, and characters in a file.

## Features

- **Count Bytes**: Calculate the size of a file in bytes.
- **Count Lines**: Count the number of lines in a file.
- **Count Words**: Count the number of words in a file.
- **Count Characters**: Count the number of characters in a file.
- **Display Version**: Show the version of the `ccwc` tool.

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)

### Clone the Repository

```bash
git clone https://github.com/yourusername/ccwc.git
cd ccwc
```

Build the Application
Copy code
```bash
go build -o ccwc
```
This will generate an executable named ccwc in your current directory.

Usage
The ccwc command provides options to count bytes, lines, words, or characters in a file.

Syntax
```bash
Copy code
ccwc [flags]
Flags
-c, --count [path]: Count the size of the file in bytes.
-l, --lines [path]: Count the number of lines in the file.
-w, --words [path]: Count the number of words in the file.
-m, --chars [path]: Count the number of characters in the file.
--help: Display help information.
version: Display the version of ccwc.
Note: Only one of the counting flags (-c, -l, -w, -m) should be used at a time.
```
Dependencies
Cobra: A library for creating powerful modern CLI applications.
