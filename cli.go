package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	dooCmd  cli.Command
	servCmd cli.Command
)

func init() {
	dooCmd = cli.Command{
		Name:  "doo",
		Usage: "doo doo doo",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "lang, l",
				Value:  "english",
				Usage:  "languages for this program",
				EnvVar: "GIN_LANG",
			},
		},
		Action: func(c *cli.Context) error {
			logrus.Info(c.String("lang"))
			return nil
		},
	}

	servCmd = cli.Command{
		Name:  "serve",
		Usage: "star this program and listenning a port",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "port, p",
				Value:  "8484",
				Usage:  "which port should be listened",
				EnvVar: "PORT",
			},
		},
		Action: func(c *cli.Context) error {
			logrus.Info(c.String("port"))
			return nil
		},
	}
}
