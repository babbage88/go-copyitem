package main

import (
	"errors"
	"fmt"
	"os"
)

type SizeInKb interface {
	GetSizeInKB() int64
}

type SizeInMb interface {
	GetSizeInMB() int64
}

type GetFileStatInfo interface {
	GetFileInfo() error
}

func (f *FileInfoExtended) GetSizeInKB() int64 {
	if !f.FileExists {
		fmt.Printf("File: %s does not exist\n", f.path)
		return 0
	}
	flength := f.FsFileInfo.Size()
	kbsize := flength / 1024

	return kbsize
}

func (f *FileInfoExtended) GetSizeInMB() int64 {
	var retVal int64
	if f.FileExists {
		flength := f.FsFileInfo.Size()
		retVal = flength / 1048576

		return retVal
	} else {
		fmt.Printf("File: %s does not exist\n", f.path)
		retVal = 0
		return retVal
	}
}

func (f *FileInfoExtended) GetFileInfo() error {
	fileinfo, err := os.Stat(f.path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f.FileExists = false
			f.IsDirectory = false
			fmt.Printf("File: %s does not exist\n", f.path)
		} else {
			fmt.Printf("Error when trying to Stat source file: %s\n", f.path)
		}
		return err
	}

	// If no error, the file exists, so set the fields
	f.FileExists = true
	f.FsFileInfo = fileinfo
	f.IsDirectory = f.FsFileInfo.IsDir()

	return nil
}
