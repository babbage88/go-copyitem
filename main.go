package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

type FileCopyJob struct {
	SourceFile      FileInfoExtended `json:"sourcefileinfo"`
	DestinationFile FileInfoExtended `json:"sourcefileinfo"`
}

type FileInfoExtended struct {
	FsFileInfo fs.FileInfo `json:"fsfileinfo"`
}

type SizeInKb interface {
	GetSizeInKB() int64
}

type SizeInMb interface {
	GetSizeInMB() int64
}

type PercentCopiedString interface {
	GetCopyProgressPercentStr() string
}

type PercentCopiedInt64 interface {
	GetCopyProgressPercentInt64() int64
}

func (f *FileInfoExtended) GetSizeInKB() int64 {
	flength := f.FsFileInfo.Size()
	kbsize := flength / 1024

	return kbsize
}

func (f *FileInfoExtended) GetSizeInMB() int64 {
	flength := f.FsFileInfo.Size()
	mbsize := flength / 1048576

	return mbsize
}

func (f *FileCopyJob) GetCopyProgressPercentStr() string {
	dstlen := f.DestinationFile.FsFileInfo.Size()
	srclength := f.SourceFile.FsFileInfo.Size()

	cursizedec := dstlen / srclength
	cursize := cursizedec * 100
	cursizestr := fmt.Sprint(cursize, "%")

	fmt.Printf("%s Percent Copied", cursizestr)

	return cursizestr

}

func (f *FileCopyJob) GetCopyProgressPercentInt64() int64 {
	dstlen := f.DestinationFile.FsFileInfo.Size()
	srclength := f.SourceFile.FsFileInfo.Size()

	cursizedec := dstlen / srclength
	cursize := cursizedec * 100

	fmt.Printf("%v Percent Copied", cursize)

	return cursize

}

func main() {
	source := flag.String("source", ".", "Source file or directory to copy")
	destination := flag.String("destination", "C:\temp", "Destination to Copy to.")
	flag.Parse()

	coloredsource := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 96, *source)
	coloreddestination := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, *destination)

	src, err := os.Stat(*source)
	if err != nil {
		fmt.Errorf("Error when trying to Stat source file %s\n", src)
	}

	dst, err := os.Stat(*destination)
	if err != nil {
		fmt.Errorf("Error when trying to Stat source file %s\n", dst)
	}
	var srcfileinfo FileInfoExtended
	var dstfileinfo FileInfoExtended
	var filecopyjob FileCopyJob

	srcfileinfo.FsFileInfo = src
	dstfileinfo.FsFileInfo = dst
	filecopyjob.SourceFile = srcfileinfo
	filecopyjob.DestinationFile = dstfileinfo

	sizemb := filecopyjob.SourceFile.GetSizeInKB()
	isSrcDir := filecopyjob.SourceFile.FsFileInfo.IsDir()
	isDstDir := filecopyjob.DestinationFile.FsFileInfo.IsDir()

	fmt.Printf("sizemb of %s is %v\n", coloredsource, sizemb)

	if isSrcDir {
		fmt.Printf("The source file specified: %s is a Directory\n", coloredsource)
	} else {
		fmt.Printf("The source file specified: %s is not a Directory.\n", coloredsource)
	}

	if isDstDir {
		fmt.Printf("The destination file specified: %s is a Directory\n", coloreddestination)
	} else {
		fmt.Printf("The destination file specified: %s is not a Directory.\n", coloreddestination)
	}

}
