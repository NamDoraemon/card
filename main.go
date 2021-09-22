package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func startApp(ctx *cli.Context) error {
	fmt.Println("CLI Start")
	return nil
}

func createTableDb(ctx *cli.Context) error {
	fmt.Println("CLI CreateTableDB")
	return nil
}

func setupCli() error {
	appCli := cli.NewApp()

	appCli.Action = func(c *cli.Context) error {
		return errors.New("Wow, ^.^ dumb")
	}

	appCli.Commands = []*cli.Command{
		{Name: "start", Action: startApp},
		{Name: "createDb", Action: createTableDb},
	}

	return appCli.Run(os.Args)
}

func main() {
	log.Print("Start Project")
	if err := setupCli(); err != nil {
		panic(err)
	}
}
