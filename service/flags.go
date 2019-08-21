package service

import (
	"github.com/namsral/flag"
	"github.com/pkg/errors"
)

type (
	configuration struct {
		mode   string
		output string

		sourcePath string

		http struct {
			addr    string
			logging bool
			pretty  bool
			tracing bool
		}
	}
)

var config *configuration

func (c *configuration) Validate() error {
	if c == nil {
		return errors.New("Config is not initialized, need to call Flags()")
	}
	if c.http.addr == "" {
		return errors.New("No HTTP Addr is set, can't listen for HTTP")
	}
	return nil
}

func Flags(prefix ...string) {
	if config != nil {
		return
	}
	config = new(configuration)

	p := func(s string) string {
		if len(prefix) > 1 {
			return prefix[0] + "-" + s
		}
		return s
	}

	flag.StringVar(&config.http.addr, p("http-addr"), ":3000", "Listen address for HTTP server")
	flag.BoolVar(&config.http.logging, p("http-log"), true, "Enable/disable HTTP request log")
	flag.BoolVar(&config.http.pretty, p("http-pretty-json"), false, "Prettify returned JSON output")
	flag.BoolVar(&config.http.tracing, p("http-error-tracing"), false, "Return error stack frame")

	flag.StringVar(&config.mode, "mode", "cli", "Runtime mode for smartcrop")
	flag.StringVar(&config.output, "output", "", "Output file (empty = standard out)")
	flag.StringVar(&config.sourcePath, "sourcePath", ".", "Source path for images")
	flag.Parse()
}
