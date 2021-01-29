package main

import (
	"flag"
	"fmt"

	"point/cmd/server/config"

	"point/internal/core"
	"point/internal/pkg/hd"
	"point/internal/server"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// @title 优药积分服务
// @version 1.0
// @description 优药积分系统 api
// @contact.name wenlong
// @contact.url http://yy-git.youyao99.com/youyao/point
// @contact.email wenlong.chen@youyaomedtech.com
// @host http://point:8080
func main() {
	var envfile string
	flag.StringVar(&envfile, "env-file", ".env", "Read in a file of environment variables")
	flag.Parse()

	godotenv.Load(envfile)
	config, err := config.Environ()
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: invalid configuration")
	}

	initLogging(config)

	// if trace level logging is enabled, output the
	// configuration parameters.
	if logrus.IsLevelEnabled(logrus.TraceLevel) {
		fmt.Println(config.String())
	}

	app, err := InitializeApplication(config)
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: cannot initialize server")
	}

	g := errgroup.Group{}
	g.Go(func() error {
		logrus.WithFields(
			logrus.Fields{
				"proto": config.Server.Proto,
				"host":  config.Server.Host,
				"port":  config.Server.Port,
				"url":   config.Server.Addr,
			},
		).Infoln("starting the http server")
		return app.server.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		logrus.WithError(err).Fatalln("program terminated")
	}
}

// helper function configures the logging.
func initLogging(c config.Config) {
	if c.Logging.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if c.Logging.Trace {
		logrus.SetLevel(logrus.TraceLevel)
	}
	if c.Logging.Text {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   c.Logging.Color,
			DisableColors: !c.Logging.Color,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: c.Logging.Pretty,
		})
	}
}

// application is the main struct for the point server.
type application struct {
	server  *server.Server
	assets  core.UserAssetsStore
	hashids *hd.HD
}

// newApplication creates a new application struct.
func newApplication(
	server *server.Server,
	assets core.UserAssetsStore,
	hashids *hd.HD,
) application {
	return application{
		assets:  assets,
		server:  server,
		hashids: hashids,
	}
}
