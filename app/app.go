package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/peterbourgon/ff"
)

func CLI(args []string) error {
	fl := flag.NewFlagSet("app", flag.ContinueOnError)
	silent := fl.Bool("quiet", false, "don't log output")
	fl.Usage = func() {
		fmt.Fprintf(fl.Output(), `slumber - Like Unix sleep but takes minutes, hours, etc.

Usage:

	slumber [options] <wake time>

Wake time may be a duration (e.g. "1h2m3s") or a target time (e.g. "1:00pm" or "13:02:03").

Options:
`)
		fl.PrintDefaults()
	}
	if err := ff.Parse(fl, args, ff.WithEnvVarPrefix("SLEEP_FOR")); err != nil {
		return err
	}

	if fl.NArg() != 1 {
		fl.Usage()
		return flag.ErrHelp
	}

	arg := fl.Arg(0)
	duration, err := parseTime(arg)
	if err != nil {
		fmt.Fprintf(fl.Output(), "bad argument %q: %v\n", arg, err)
		fl.Usage()
		return flag.ErrHelp
	}

	return appExec(duration, !*silent)
}

func parseTime(s string) (d time.Duration, err error) {
	for _, format := range []string{
		time.Kitchen, strings.ToLower(time.Kitchen), "15:04", "15:04:05",
	} {
		t, err := time.Parse(format, s)
		if err == nil {
			now := time.Now()
			t = time.Date(
				now.Year(), now.Month(), now.Day(),
				t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
				time.Local)
			if !t.After(now) {
				t = t.AddDate(0, 0, 1)
			}
			return t.Sub(now), nil
		}
	}

	d, err = time.ParseDuration(s)
	return
}

func appExec(duration time.Duration, verbose bool) error {
	l := nooplogger
	if verbose {
		l = log.New(os.Stderr, "slumber ", log.LstdFlags).Printf
	}
	a := app{duration, l}
	if err := a.exec(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return err
	}
	return nil
}

type logger = func(format string, v ...interface{})

func nooplogger(format string, v ...interface{}) {}

type app struct {
	duration time.Duration
	log      logger
}

func (a *app) exec() (err error) {
	a.log("starting slumbering for %v", a.duration)
	defer func() { a.log("done") }()
	time.Sleep(a.duration)
	return err
}
