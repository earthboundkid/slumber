# Slumber [![GoDoc](https://godoc.org/github.com/carlmjohnson/slumber?status.svg)](https://godoc.org/github.com/carlmjohnson/slumber) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/slumber)](https://goreportcard.com/report/github.com/carlmjohnson/slumber)

Like Unix sleep but takes minutes, hours, etc.

## Installation

First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN="$(pwd)" GOPATH="$(mktemp -d)" go get github.com/carlmjohnson/slumber
```

## Screenshots

```bash
$ slumber -h
slumber - Like Unix sleep but takes minutes, hours, etc.

Usage:

        slumber [options] <wake time>

Wake time may be a duration (e.g. "1h2m3s") or a target time (e.g. "1:00pm" or "13:02:03").

Options:
  -quiet
        don't log output

$ slumber 1s
slumber 2019/10/23 12:11:04 starting sleeping for 1s
slumber 2019/10/23 12:11:05 done

$ slumber 12:11:30
slumber 2019/10/23 12:11:16 starting sleeping for 13.230717s
slumber 2019/10/23 12:11:30 done
```
