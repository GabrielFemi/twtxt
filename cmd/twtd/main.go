package main

import (
	"fmt"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/prologic/twtxt"
)

var (
	bind    string
	debug   bool
	version bool

	data     string
	store    string
	name     string
	register bool
	baseURL  string
)

func init() {
	flag.BoolVarP(&version, "version", "v", false, "display version information")
	flag.BoolVarP(&debug, "debug", "D", false, "enable debug logging")
	flag.StringVarP(&bind, "bind", "b", "0.0.0.0:8000", "[int]:<port> to bind to")

	flag.StringVarP(&data, "data", "d", twtxt.DefaultData, "data directory")
	flag.StringVarP(&store, "store", "s", twtxt.DefaultStore, "store to use")
	flag.StringVarP(&name, "name", "n", twtxt.DefaultName, "set the instance's name")
	flag.BoolVarP(&register, "register", "r", twtxt.DefaultRegister, "enable user registration")
	flag.StringVarP(&baseURL, "base-url", "u", twtxt.DefaultBaseURL, "base url to use for app")
}

func main() {
	flag.Parse()

	if version {
		fmt.Printf("twtxt v%s", twtxt.FullVersion())
		os.Exit(0)
	}

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	svr, err := twtxt.NewServer(bind,
		twtxt.WithData(data),
		twtxt.WithName(name),
		twtxt.WithStore(store),
		twtxt.WithBaseURL(baseURL),
		twtxt.WithRegister(register),
	)
	if err != nil {
		log.WithError(err).Fatal("error creating server")
	}

	log.Infof("%s listening on http://%s", path.Base(os.Args[0]), bind)
	if err := svr.Run(); err != nil {
		log.WithError(err).Fatal("error running or shutting down server")
	}
}
