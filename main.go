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
	destination := flag.String("destination", "C:\tmep", "Destination to Copy to.")
	flag.Parse()

	fmt.Printf("Copying %s to %s", *source, *destination)

	src, err := os.Stat(*source)
	if err != nil {
		fmt.Errorf("Error when trying to Stat source file %s\n", src)
	}

	srcsize := src.Size()
	kbsize := srcsize / 1024

	fmt.Printf("Size of source file %s is %v in bytes, or %v in KiloBytes\n", src.Name(), srcsize, kbsize)

}
