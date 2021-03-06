package main

import (
  "time"
	"os"

  "github.com/joho/godotenv"
  "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Loading .env error, now continue starting program with default value...")
	}

	app := cli.NewApp()
	app.Name = "gin boilerplate"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "leopku",
			Email: "leo@himysql.com",
		},
	}
	app.Copyright = "(c) 2018 leopku"
	app.Commands = []cli.Command{
		dooCmd,
    migrateCmd,
    servCmd,
	}

	app.Run(os.Args)
}
