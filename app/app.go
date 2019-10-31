package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/carlmjohnson/flagext"
	"github.com/peterbourgon/ff"
)

func CLI(args []string) error {
	a, err := parseArgs(args)
	if err != nil {
		return err
	}
	if err := a.exec(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return err
	}
	return nil
}

func parseArgs(args []string) (*app, error) {
	fl := flag.NewFlagSet("app", flag.ContinueOnError)
	l := log.New(nil, "slumber ", log.LstdFlags)
	fl.Var(
		flagext.Logger(l, flagext.LogSilent),
		"quiet",
		`don't log output`,
	)

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
		return nil, err
	}

	if fl.NArg() != 1 {
		fl.Usage()
		return nil, flag.ErrHelp
	}

	arg := fl.Arg(0)
	target, err := parseTime(arg)
	if err != nil {
		fmt.Fprintf(fl.Output(), "bad argument %q: %v\n", arg, err)
		fl.Usage()
		return nil, flag.ErrHelp
	}
	return &app{target, l}, nil
}

func parseTime(s string) (t time.Time, err error) {
	// Do this first so less time passes
	if d, err := time.ParseDuration(s); err == nil {
		return time.Now().Add(d), nil
	}

	for _, format := range []string{
		time.Kitchen, strings.ToLower(time.Kitchen), "15:04", "15:04:05",
	} {
		t, err = time.Parse(format, s)
		if err == nil {
			now := time.Now()
			t = time.Date(
				now.Year(), now.Month(), now.Day(),
				t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
				time.Local)
			if !t.After(now) {
				t = t.AddDate(0, 0, 1)
			}
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse wake time: %q", s)
}

type app struct {
	target time.Time
	*log.Logger
}

func (a *app) exec() (err error) {
	total := time.Until(a.target)
	a.Printf("starting slumbering until %s (%v)",
		a.target.Format(time.Stamp), total)
	defer func() { a.Print("done") }()

	const fraction = 10
	delta := total / fraction
	for {
		if until := time.Until(a.target); until < delta {
			if until < 1 {
				break
			}
			time.Sleep(time.Until(a.target))
			break
		}
		time.Sleep(delta)
	}
	return err
}
