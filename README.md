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

- [PortSwigger - File path traversal, simple case](https://portswigger.net/web-security/file-path-traversal/lab-simple)
- [PortSwigger - File path traversal, traversal sequences blocked with absolute path bypass](https://portswigger.net/web-security/file-path-traversal/lab-absolute-path-bypass)
- [PortSwigger - File path traversal, traversal sequences stripped non-recursively](https://portswigger.net/web-security/file-path-traversal/lab-sequences-stripped-non-recursively)
- [PortSwigger - File path traversal, traversal sequences stripped with superfluous URL-decode](https://portswigger.net/web-security/file-path-traversal/lab-superfluous-url-decode)
- [PortSwigger - File path traversal, validation of start of path](https://portswigger.net/web-security/file-path-traversal/lab-validate-start-of-path)
- [PortSwigger - File path traversal, validation of file extension with null byte bypass](https://portswigger.net/web-security/file-path-traversal/lab-validate-file-extension-null-byte-bypass)
