package main

import (
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/urfave/cli/v2"
)

func testDev(src string, dst string, w int) {
	if runtime.GOOS == "windows" {
		log.Println("Hello from Windows")
	}
	srcPath, srcFile := filepath.Split(src)
	dstPath, dstFile := filepath.Split(dst)
	log.Printf("src: %s srcDir: %s srcFile: %s\n", src, srcPath, srcFile)
	log.Printf("dst: %s dstDir: %s dstFile: %s\n", dst, dstPath, dstFile)
	progressBarConfig := NewProgressBarConfig(WithProgressBarWidth(w))

	filecopyjob := NewFileCopyJob(WithSourceFilePath(src), WithDestinationFilePath(dst), WithProgressBarConfig(progressBarConfig))

	if filecopyjob.DestinationFile.IsDirectory {
		if !filecopyjob.SourceFile.IsDirectory {
			log.Printf("src is a File: %s\n", src)
			filecopyjob.DestinationFile.path = filepath.Join(dst, srcFile)
			testC := filecopyjob.PrettyPrintDst()
			log.Printf("New dst: %s\n", testC)

		}
		log.Printf("The specified dst: %s  is a directory\n", dst)
		return
	}
}

func (f *FileCopyJob) ParsePathParams() error {
	_, srcFile := filepath.Split(f.SourceFile.path)

	if f.DestinationFile.IsDirectory {
		if !f.SourceFile.IsDirectory {
			f.DestinationFile.path = filepath.Join(f.DestinationFile.path, srcFile)

			log.Printf("No filename specified for destination. Using Source filename: %s\n", f.DrawColoredString(srcFile, 96))
			log.Printf("No filename specified for destination. Using Source filename: %s\n", f.PrettyPrintDst())

		}

		return nil
	}
	return nil
}

func (f FileInfoExtended) PrettyPrintSizeString() string {
	size := int(f.SizeBytes)

	if size < 1024 {
		return f.PrettyStringSizeBytes()
	} else if size >= 1<<10 {
		return f.PrettyStringSizeKB()
	} else if size < 1048576 {
		return f.PrettyStringSizeMB()
	} else if size < 1073741824 {
		return f.PrettyStringSizeGB()
	}

	return f.PrettyStringSizeBytes()
}

func cmdCopyFileJob(src string, dst string, w int) {

	progressBarConfig := NewProgressBarConfig(WithProgressBarWidth(w))

	filecopyjob := NewFileCopyJob(WithSourceFilePath(src), WithDestinationFilePath(dst), WithProgressBarConfig(progressBarConfig))

	filecopyjob.ParsePathParams()

	filecopyjob.Start()
}

func CopyJobCommand() (appInst *cli.App) {
	appInst = &cli.App{
		Name:                 "gocp",
		Version:              "1.0.2",
		Compiled:             time.Now(),
		Args:                 true,
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
			&cli.BoolFlag{
				Name:    "test",
				Aliases: []string{"tst"},
				Value:   false,
				Usage:   "Used for testing/development",
			},
		},
		Action: func(cCtx *cli.Context) (err error) {
			if cCtx.Bool("test") {
				testDev(cCtx.String("source"), cCtx.String("destination"), cCtx.Int("width"))
				return nil
			}
			if cCtx.NArg() == 0 {
				cmdCopyFileJob(cCtx.String("source"), cCtx.String("destination"), cCtx.Int("width"))
				return nil
			}
			log.Printf("args: %+v", cCtx.Args())

			cmdCopyFileJob(cCtx.Args().Get(0), cCtx.Args().Get(1), cCtx.Int("width"))
			return nil
		},
	}
	return appInst
}
