package main

import (
	"fmt"
	"io"
	"os"
)

type FileCopyJob struct {
	SourceFile      FileInfoExtended `json:"sourcefileinfo"`
	DestinationFile FileInfoExtended `json:"sourcefileinfo"`
}

type IFileCopyJob interface {
	GetCopyProgressPercentStr() string
	GetCopyProgressPercentInt64() int64
	PrettyPrintSrc() string
	PrettyPrintDst() string
	CopyFile() error
}

func (f *FileCopyJob) PrettyPrintSrc() string {
	coloredsource := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 96, f.SourceFile.path)
	return coloredsource
}

func (f *FileCopyJob) PrettyPrintDst() string {
	colordestination := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 92, f.DestinationFile.path)
	return colordestination
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

func (f *FileCopyJob) CopyFile() error {
	fmt.Printf("Starting File Copy Job src: %s\ndst: %s\nsize_kb: %s", f.SourceFile.path, f.DestinationFile.path, f.SourceFile.GetSizeInKB())
	src, err := os.Open(f.SourceFile.path)
	if err != nil {
		fmt.Printf("Error Opening source file %s\n", f.PrettyPrintSrc())
		return err
	}
	defer src.Close()

	newfile, err := os.Create(f.DestinationFile.path)
	if err != nil {
		fmt.Printf("Error Creating new destination file: %s \n", f.PrettyPrintDst())
		return err
	}

	defer func() {
		// return any error when closing the new file, but only if there are none preceeding it.
		if crt := newfile.Close(); err == nil {
			err = crt
		}
	}()

	_, err = io.Copy(newfile, src)
	if err != nil {
		fmt.Printf("The copy of %s to %s failed\n", f.PrettyPrintSrc(), f.PrettyPrintDst())
	}

	return err
}
