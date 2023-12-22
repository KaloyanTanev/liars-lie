package main

import (
	"log"
	"os"

	myCLI "liars-lie/app/cli"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Liars Lie",
		Usage: "Everyone's favourite game",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start the agents",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "value", Usage: "network value", Required: true},
					&cli.IntFlag{Name: "max-value", Usage: "max value for liars", Required: true},
					&cli.IntFlag{Name: "num-agents", Usage: "total number of agents", Required: true},
					&cli.Float64Flag{Name: "liar-ratio", Usage: "ratio of liars", Required: true},
				},
				Action: func(cCtx *cli.Context) error {
					myCLI.Start(cCtx.Int("value"), cCtx.Int("max-value"), cCtx.Int("num-agents"), cCtx.Float64("liar-ratio"))
					return nil
				},
			},
			{
				Name:  "play",
				Usage: "play round",
				Action: func(cCtx *cli.Context) error {
					myCLI.Play()
					return nil
				},
			},
			{
				Name:  "stop",
				Usage: "stop all agents",
				Action: func(cCtx *cli.Context) error {
					myCLI.Stop()
					return nil
				},
			},
			{
				Name:  "extend",
				Usage: "add new agents",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "value", Usage: "network value", Required: true},
					&cli.IntFlag{Name: "max-value", Usage: "max value for liars", Required: true},
					&cli.IntFlag{Name: "num-agents", Usage: "total number of agents", Required: true},
					&cli.Float64Flag{Name: "liar-ratio", Usage: "ratio of liars", Required: true},
				},
				Action: func(cCtx *cli.Context) error {
					myCLI.Extend(cCtx.Int("value"), cCtx.Int("max-value"), cCtx.Int("num-agents"), cCtx.Float64("liar-ratio"))
					return nil
				},
			},
			{
				Name:  "playexpert",
				Usage: "play in expert mode",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "num-agents", Usage: "number of agents queried", Required: true},
					&cli.Float64Flag{Name: "liar-ratio", Usage: "suspected ratio of liers", Required: true},
				},
				Action: func(cCtx *cli.Context) error {
					myCLI.PlayExpert(cCtx.Int("num-agents"), cCtx.Float64("liar-ratio"))
					return nil
				},
			},
			{
				Name:  "kill",
				Usage: "kill an agent with id",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "id", Usage: "id of an agent", Required: true},
				},
				Action: func(cCtx *cli.Context) error {
					myCLI.Kill(cCtx.Int("id"))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
