package main

import (
  "os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gin boilerplate"
	app.Version = "0.0.1"
	app.Complied = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name: "leopku",
			Email: "leo@himysql.com",
		},
  }
  app.Copyright = "(c) 2018 leopku"
  app.Run(os.Args)
}
