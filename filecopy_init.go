package main

type FileCopyJobOptions func(*FileCopyJob)
type FileIntoExtendedOptions func(*FileInfoExtended)

func WithSourceFilePath(path string) FileCopyJobOptions {
	return func(copyJob *FileCopyJob) {
		var copyJobFile FileInfoExtended
		copyJobFile.path = path
		copyJobFile.GetFileInfo()
		copyJob.SourceFile = &copyJobFile
	}
}

func WithDestinationFilePath(path string) FileCopyJobOptions {
	return func(copyJob *FileCopyJob) {
		var copyJobFile FileInfoExtended
		copyJobFile.path = path
		copyJobFile.GetFileInfo()
		copyJob.DestinationFile = &copyJobFile
	}
}

func WithSourceColor(i int) FileCopyJobOptions {
	return func(copyJob *FileCopyJob) {
		copyJob.SrcColor = i
	}
}

func WithDestinationColor(i int) FileCopyJobOptions {
	return func(copyJob *FileCopyJob) {
		copyJob.DstColor = i
	}
}

func WithSourceFile(sourceFileInfo FileInfoExtended) FileCopyJobOptions {
	return func(f *FileCopyJob) {
		f.SourceFile = &sourceFileInfo
	}
}

func WithDestinationFile(destinationFileInfo FileInfoExtended) FileCopyJobOptions {
	return func(f *FileCopyJob) {
		f.DestinationFile = &destinationFileInfo
	}
}

func WithProgressBarConfig(p *ProgressBarConfig) FileCopyJobOptions {
	return func(f *FileCopyJob) {
		f.ProgressBarConfig = p
	}
}

func NewFileCopyJob(opts ...FileCopyJobOptions) *FileCopyJob {
	const (
		srcColor = 96
		dstColor = 92
	)

	fileCopyJob := &FileCopyJob{SrcColor: srcColor, DstColor: dstColor}
	fileCopyJob.ProgressBarConfig = NewProgressBarConfig()

	for _, opt := range opts {
		opt(fileCopyJob)
	}
	return fileCopyJob
}
