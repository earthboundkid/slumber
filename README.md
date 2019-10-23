# $NAME [![GoDoc](https://godoc.org/github.com/carlmjohnson/sleep-for?status.svg)](https://godoc.org/github.com/carlmjohnson/sleep-for) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/sleep-for)](https://goreportcard.com/report/github.com/carlmjohnson/sleep-for)

Like Unix sleep but takes minutes, hours, etc.

## Installation

First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN="$(pwd)" GOPATH="$(mktemp -d)" go get github.com/carlmjohnson/sleep-for
```

## Screenshots

```bash
$ sleep-for -h
sleep-for - Like Unix sleep but takes minutes, hours, etc.

Usage:

        sleep-for [options]

Options:
  -duration duration
        how long to sleep (default 1s)
  -verbose
        log debug output

$ sleep-for . -verbose
sleep-for 2019/10/22 23:42:18 starting sleeping for 1s
sleep-for 2019/10/22 23:42:19 done
```
