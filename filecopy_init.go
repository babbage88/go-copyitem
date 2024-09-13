package main

import "fmt"

type FileCopyJobOptions func(*FileCopyJob)
type FileIntoExtendedOptions func(*FileInfoExtended)
type ProgressBarConfigOptions func(*ProgressBarConfig)

func WithProgressBarWidth(width int) ProgressBarConfigOptions {
	return func(p *ProgressBarConfig) {
		p.Width = width
	}
}

func WithProgressFillCharacter(s string) ProgressBarConfigOptions {
	return func(p *ProgressBarConfig) {
		p.FillCharacter = s
	}
}

func WithProgressRemaingCharacter(s string) ProgressBarConfigOptions {
	return func(p *ProgressBarConfig) {
		p.RemainingCharacter = s
	}
}

func (f *ProgressBarConfig) DrawColoredString(s string, color int) string {
	coloredString := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, s)
	return coloredString
}

func NewProgressBarConfig(opts ...ProgressBarConfigOptions) *ProgressBarConfig {
	const (
		width = 50
	)
	progBarConf := &ProgressBarConfig{Width: width}
	progBarConf.FillCharacter = progBarConf.DrawColoredString("#", 92)
	progBarConf.RemainingCharacter = progBarConf.DrawColoredString("-", 96)

	for _, opt := range opts {

		opt(progBarConf)

	}

	return progBarConf
}

func WithSourceFilePath(path string) FileCopyJobOptions {
	return func(copyJob *FileCopyJob) {
		var copyJobFile FileInfoExtended
		copyJobFile.path = path
		copyJobFile.GetFileInfo()
		copyJob.SourceFile = copyJobFile
	}
}

func WithDestinationFilePath(path string) FileCopyJobOptions {
	return func(copyJob *FileCopyJob) {
		var copyJobFile FileInfoExtended
		copyJobFile.path = path
		copyJobFile.GetFileInfo()
		copyJob.DestinationFile = copyJobFile
	}
}

func WithSourceFile(sourceFileInfo FileInfoExtended) FileCopyJobOptions {
	return func(f *FileCopyJob) {
		f.SourceFile = sourceFileInfo
	}
}

func WithDestinationFile(destinationFileInfo FileInfoExtended) FileCopyJobOptions {
	return func(f *FileCopyJob) {
		f.DestinationFile = destinationFileInfo
	}
}

func WithProgressBarConfig(p *ProgressBarConfig) FileCopyJobOptions {
	return func(f *FileCopyJob) {
		f.ProgressBarConfig = p
	}
}

func NewFileCopyJob(opts ...FileCopyJobOptions) *FileCopyJob {
	fileCopyJob := &FileCopyJob{}
	fileCopyJob.ProgressBarConfig = NewProgressBarConfig()

	for _, opt := range opts {
		opt(fileCopyJob)
	}
	return fileCopyJob
}
