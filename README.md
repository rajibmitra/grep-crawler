# Go Grep Tool

A Go command-line tool that searches for a specified pattern in directories, excluding the `.git` directory.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

To use the Go Grep Tool, you need to have Go installed on your system.

1. Clone this repository or download the source code.

```shell
git clone https://github.com/rajibmitra/grep-crawler.git

cd grep-crawler

go build . 

```

## Usage

The Go Grep Tool allows you to search for a specific pattern in directories while excluding the .git directory. Here's the basic usage:

```shell

./grep-crawler <pattern> <directory1> [directory2...]

```

<pattern>: The pattern you want to search for.

<directory1> [directory2...]: One or more directories where you want to search.

## Examples

Search for the pattern "searchme" in the "project" directory:

## Contributing

Contributions are welcome! If you find a bug or have an enhancement in mind, please open an issue or submit a pull request.

## License 

This project is licensed under the MIT License.