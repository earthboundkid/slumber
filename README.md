# Sleep-for [![GoDoc](https://godoc.org/github.com/carlmjohnson/sleep-for?status.svg)](https://godoc.org/github.com/carlmjohnson/sleep-for) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/sleep-for)](https://goreportcard.com/report/github.com/carlmjohnson/sleep-for)

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

        sleep-for [options] <wake time>

Wake time may be a duration (e.g. "1h2m3s") or a target time (e.g. "1:00pm" or "13:02:03").

Options:
  -quiet
        don't log output

$ sleep-for 1s
sleep-for 2019/10/23 12:11:04 starting sleeping for 1s
sleep-for 2019/10/23 12:11:05 done

$ sleep-for 12:11:30
sleep-for 2019/10/23 12:11:16 starting sleeping for 13.230717s
sleep-for 2019/10/23 12:11:30 done
```
