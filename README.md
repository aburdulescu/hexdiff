# hexdiff
Diff two hex strings

## Install
- [install go](https://golang.org/doc/install)
- run `go install bandr.me/p/hexdiff@latest`

## Usage
Run the tool with two args: first hex string and second hex string.

E.g.
``` bash
hexdiff 3024a7c5ecb289adcaa00a06082a8648 3002e3646fc42be9f24c6dcaa1b43b29
```

It will print the two args, each on its own line, with the parts that differ colored in red.
