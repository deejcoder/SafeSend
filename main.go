package main

import (
	"SafeSend/server/config"
	"SafeSend/server/storage"
	"SafeSend/server/user/repository"
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
	db.Connect()

	repository.CreateUser(db, "deejcoder2")
	users, err := repository.GetUsers(db)
	if err != nil {
		log.Error(err)
	}

	for _, user := range users {
		fmt.Printf("User %d: %s", user.ID, user.DisplayName)
	}

}
