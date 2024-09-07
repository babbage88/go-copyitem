package main

import (
	"fmt"
	"io"
	"os"
)

type PercentCopiedString interface {
	GetCopyProgressPercentStr() string
}

type PercentCopiedInt64 interface {
	GetCopyProgressPercentInt64() int64
}

type CopyFileSimple interface {
	CopyFile() error
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
	src, err := os.Open(f.SourceFile.path)
	if err != nil {
		return err
	}
	defer src.Close()

	newfile, err := os.Create(f.DestinationFile.path)
	if err != nil {
		return err
	}

	defer func() {
		// return any error when closing the new file, but only if there are none preceeding it.
		if crt := newfile.Close(); err == nil {
			err = crt
		}
	}()

	_, err = io.Copy(newfile, src)

	return err
}
