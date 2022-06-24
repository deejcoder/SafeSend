package main

import (
	"SafeSend/config"
	"SafeSend/storage/collections"
	storage "SafeSend/storage/database"
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

	collections.CreateUser(db, "deejcoder")
	users, err := collections.GetUsers(db)
	if err != nil {
		log.Error(err)
	}

	for idx, user := range users {
		fmt.Printf("User %d: %s", idx, user.DisplayName)
	}

	//Close()

}

func Close() {
	storage.Close()
}
