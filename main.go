package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var logger = logrus.New()

var logLevelMap = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
}

/*
type roundTripper func(*http.Request) (*http.Response, error)

func (f roundTripper) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

func reverseProxy(target string) (gin.HandlerFunc, error) {
	logger.WithField("target", target).Info("proxy")
	url, err := url.Parse(target)
	if err != nil {
		return nil, errors.Wrapf(err, "Fail to parse endpoint URL: %v", target)
	}

	requestHandler := func(req *http.Request) (*http.Response, error) {
		req.Host = url.Host
		return http.DefaultTransport.RoundTrip(req)
	}

	proxy := &httputil.ReverseProxy{
		Transport: roundTripper(requestHandler),
		Director: func(req *http.Request) {
			req.URL.Host = url.Host
			req.URL.Scheme = url.Scheme
			req.URL.Path = url.Path + req.URL.Path
		},
	}

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}, nil
}
*/

type arguments struct {
	LogLevel       string
	BindAddress    string
	BindPort       int
	StaticContents string
}

func runServer(args arguments) error {
	level, ok := logLevelMap[args.LogLevel]
	if !ok {
		return fmt.Errorf("Invalid log level: %s", args.LogLevel)
	}
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"args": args,
	}).Info("Given options")

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile(args.StaticContents, false)))
	r.GET("/api/v1/hello", func(c *gin.Context) {
		c.String(200, `{"message":"hello"}`)
	})

	if err := r.Run(fmt.Sprintf("%s:%d", args.BindAddress, args.BindPort)); err != nil {
		return err
	}

	return nil
}

func main() {
	var args arguments

	app := cli.NewApp()
	app.Name = "web-app-go"
	app.Usage = "Sample Web App"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Masayoshi Mizutani",
			Email: "mizutani@sfc.wide.ad.jp",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "log-level, l", Value: "info",
			Usage:       "Log level [trace,debug,info,warn,error]",
			Destination: &args.LogLevel,
		},
		cli.StringFlag{
			Name: "addr, a", Value: "127.0.0.1",
			Usage:       "Bind address",
			Destination: &args.BindAddress,
		},
		cli.IntFlag{
			Name: "port, p", Value: 9080,
			Usage:       "Bind port",
			Destination: &args.BindPort,
		},
		cli.StringFlag{
			Name: "static, s", Value: "./static/",
			Usage:       "Static contents path",
			Destination: &args.StaticContents,
		},
	}

	app.Action = func(c *cli.Context) error {
		if err := runServer(args); err != nil {
			return err
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logger.WithError(err).Fatal("Fatal Error")
	}
}
