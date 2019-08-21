package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/SentimensRG/sigctx"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"github.com/titpetric/factory/resputil"
)

func Start() error {
	Flags()

	// validate configuration
	if err := config.Validate(); err != nil {
		return err
	}

	// configure resputil options
	resputil.SetConfig(resputil.Options{
		Pretty: config.http.pretty,
		Trace:  config.http.tracing,
		Logger: func(err error) {
			log.Printf("Error from request: %+v", err)
		},
	})

	// run based on selected mode (cli, http)
	switch config.mode {
	case "http":
		var ctx = sigctx.New()

		log.Println("Starting http server on address " + config.http.addr)
		listener, err := net.Listen("tcp", config.http.addr)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Can't listen on addr %s", config.http.addr))
		}

		r := chi.NewRouter()

		// mount routes
		mountRoutes(r, config)

		go http.Serve(listener, r)
		<-ctx.Done()

		return nil
	case "cli":
		if len(os.Args) < 2 {
			return errors.New("Missing parameter: image file")
		}
		filename := os.Args[len(os.Args)-1]

		results, err := smartCrop(filename)
		if err != nil {
			return err
		}

		result, err := json.Marshal(results)
		if err != nil {
			return err
		}

		switch true {
		case config.output != "":
			if err := ioutil.WriteFile(config.output, result, 0644); err != nil {
				return err
			}
		default:
			fmt.Printf("%s\n", result)
		}
	default:
		return errors.New("Unknown mode: " + config.mode + ", supported are [cli, http], use -? for options.")
	}
	return nil
}
