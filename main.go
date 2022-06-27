package main

import (
	"SafeSend/pkg/config"
	"SafeSend/pkg/storage"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	app := &cli.App{
		Name: "start",
		Action: func(c *cli.Context) error {
			Start()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%v", err)
	}

}

func Start() {

	config.InitConfig()

	db := new(storage.Database)
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

}
