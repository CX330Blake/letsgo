# letsgo - Local Enumeration & Traversal Scanning in Go

# Installation && Update

> Prerequisites: You need to have Go installed and add `$GOPATH/bin` to `$PATH`, you can use `go version` to check whether you have installed it. For the `$PATH` issue, you can see the details at the [official docs](https://go.dev/doc/install#)

Install and update are quite easy, they use the same command, just run the following command then everything's done.

```bash
go install github.com/CX330Blake/letsgo@latest
```

# Usage

Use `letsgo --help` to see the full usage.

## Basic usage

Use the default settings to scan a URL.

```bash
letsgo --url <https://example.com>
```

# Labs

If you want to check what it can does, you can use the labs provided by Port Swigger. There're 6 labs which varies in defferent bypass skill to exploit the path traversal vulnerability, and **letsgo** can deal with ALL of them, quick, and precise. Use the lab to learn more about path traversal, GLHF!
