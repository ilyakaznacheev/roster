package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/ilyakaznacheev/roster"
	"github.com/ilyakaznacheev/roster/internal/config"
)

type Args struct {
	ConfigPath string
}

func main() {
	var cfg config.Application

	args := ProcessArgs(&cfg)

	// read configuration from the file and environment variables
	if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
		exitWithError(err)
	}

	if err := roster.Run(cfg); err != nil {
		exitWithError(err)
	}
}

// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg interface{}) Args {
	var a Args

	// handle command-line parameters
	f := flag.NewFlagSet("Example server", 1)
	f.StringVar(&a.ConfigPath, "c", "configs/config.yml", "Path to configuration file")

	// set custom help text
	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])
	return a
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
