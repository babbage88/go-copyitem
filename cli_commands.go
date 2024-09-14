package main

import (
	"log"
	"time"

	"github.com/urfave/cli/v2"
)

func cmdCopyFileJob(src string, dst string, w int) {
	progressBarConfig := NewProgressBarConfig(WithProgressBarWidth(w))

	filecopyjob := NewFileCopyJob(WithSourceFilePath(src), WithDestinationFilePath(dst), WithProgressBarConfig(progressBarConfig))

	filecopyjob.Start()
}

func CopyJobCommand() (appInst *cli.App) {
	appInst = &cli.App{
		Name:                 "gocp",
		Version:              "1.0.2",
		Compiled:             time.Now(),
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Justin Trahan",
				Email: "justin@trahan.dev",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "source",
				Aliases: []string{"s"},
				Usage:   "The source file to be copied",
			},
			&cli.StringFlag{
				Name:    "destination",
				Aliases: []string{"d"},
				Usage:   "The destination file to be copied",
			},
			&cli.IntFlag{
				Name:    "width",
				Aliases: []string{"w"},
				Value:   75,
				Usage:   "Width of the Progress Bar that gets drawn",
			},
		},
		Action: func(cCtx *cli.Context) (err error) {
			log.Println("Starting CopyFileJob Command")
			if cCtx.NArg() == 0 {
				log.Fatal("No CLI arguments detected!\n Please specify source and destination")
			}
			log.Printf("args: %+v", cCtx.Args())

			cmdCopyFileJob(cCtx.String("source"), cCtx.String("destination"), cCtx.Int("width"))
			return nil
		},
	}
	return appInst
}
