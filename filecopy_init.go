package main

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

func NewCopyJobProgressBarConfig(w int, fill string, remain string) ProgressBarConfig {
	progressBarConf := NewCopyJobProgressBarConfig(w, fill, remain)

	return progressBarConf
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

func WithProgressBarConfig(p ProgressBarConfig) FileCopyJobOptions {
	return func(f *FileCopyJob) {
		f.ProgressBarConfig = p
	}
}

func NewFileCopyJob(opts ...FileCopyJobOptions) *FileCopyJob {
	fileCopyJob := &FileCopyJob{}
	fillChar := fileCopyJob.DrawColoredString("#", 92)
	remChar := fileCopyJob.DrawColoredString("-", 96)
	NewCopyJobProgressBarConfig(50, fillChar, remChar)
	for _, opt := range opts {
		opt(fileCopyJob)
	}
	return fileCopyJob
}
