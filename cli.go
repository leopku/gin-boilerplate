package main

import (
  // "database/sql"

  "github.com/golang-migrate/migrate"
  _ "github.com/golang-migrate/migrate/database/postgres"
  _ "github.com/golang-migrate/migrate/source/file"
  _ "github.com/lib/pq"

  "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
  dooCmd      cli.Command
  migrateCmd  cli.Command
  servCmd     cli.Command
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

  migrateCmd = cli.Command{
    Name: "migrate",
    Usage: "migrate database",
    Action:func(c *cli.Context) error {
      // db, err := sql.Open("postgres", "postgres://gin:12345678@localhost/gin?sslmode=disable")

      m, err := migrate.New(
        "file://./migrations",
        "postgres://gin:12345678@localhost/postgres?sslmode=disable",
      )
      if err != nil {
        logrus.Error("migration error: " + err.Error())
        return err
      }
      m.Up()
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
      port := c.String("port")
			logrus.Info(port)
			return nil
		},
	}
}
