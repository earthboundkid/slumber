package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/peterbourgon/ff"
)

func CLI(args []string) error {
	fl := flag.NewFlagSet("app", flag.ContinueOnError)
	duration := fl.Duration("duration", 1*time.Second, "how long to sleep")
	verbose := fl.Bool("verbose", false, "log debug output")
	fl.Usage = func() {
		fmt.Fprintf(fl.Output(), `sleep-for - Like Unix sleep but takes minutes, hours, etc.

Usage:

	sleep-for [options]

Options:
`)
		fl.PrintDefaults()
	}
	if err := ff.Parse(fl, args, ff.WithEnvVarPrefix("SLEEP_FOR")); err != nil {
		return err
	}

	return appExec(*duration, *verbose)
}

func appExec(duration time.Duration, verbose bool) error {
	l := nooplogger
	if verbose {
		l = log.New(os.Stderr, "sleep-for ", log.LstdFlags).Printf
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
	a.log("starting")
	defer func() { a.log("done") }()
	time.Sleep(a.duration)
	return err
}
