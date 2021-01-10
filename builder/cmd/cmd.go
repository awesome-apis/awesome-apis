package main

import (
	"github.com/awesome-apis/awesome-apis/builder"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "build",
		Action: func(c *cli.Context) error {
			apisData, err := ioutil.ReadFile(c.Args().First())
			if err != nil {
				return err
			}

			var list []*builder.API
			err = yaml.Unmarshal(apisData, &list)
			if err != nil {
				return err
			}

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			return builder.Render(list, dir)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
