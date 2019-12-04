package main

import (
	"flag"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"os"
	"testing"
)

func FeatureContext(s *godog.Suite) {
	// TODO insert steps

	s.BeforeScenario(func(interface{}) {
		// TODO initialize scenario if necessary
	})
}

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("login", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
