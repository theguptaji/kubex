package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "load",
				Aliases: []string{"l"},
				Usage:   "load a project into kubex",
				Action:  loadProject,
			},
			{
				Name:    "unload",
				Aliases: []string{"u"},
				Usage:   "unload this project from kubex",
				Action:  unloadProject,
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "build the manifests for the current project",
				Action:  buildProject,
			},
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "show the network flow for the current project",
				Action:  showProject,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func loadProject(c *cli.Context) error {
	fmt.Println("Loading project:", c.Args().First())
	return nil
}

func unloadProject(c *cli.Context) error {
	fmt.Println("Unloading project:", c.Args().First())
	return nil
}

func buildProject(c *cli.Context) error {
	fmt.Println("Building project:", c.Args().First())
	return nil
}

func showProject(c *cli.Context) error {
	fmt.Println("Showing project:", c.Args().First())
	return nil
}