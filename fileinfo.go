package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type FileInfoExtended struct {
	FsFileInfo  fs.FileInfo `json:"fsfileinfo"`
	path        string      `json:"path"`
	IsDirectory bool        `json:"isDirectory"`
	FileExists  bool        `json:"files_exists"`
	SizeBytes   float64     `json:"sizeBytes"`
}

type CopyJobFile interface {
	GetFileInfo() error
	GetSizeInKB() float64
	GetSizeInMB() float64
	GetSizeInGB() float64
	GetSizeBytes() float64
	CheckSize() float64
	PrettyPrintSizeBytes() string
	PrettyPrintSizeKB() string
	PrettyPrintSizeMB() string
	PrettyPrintSizeGB() string
}

func (f *FileCopyJob) PrettyPrintSizeBytes() string {
	coloredsource := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 96, f.DestinationFile.SizeBytes)
	return coloredsource
}

func (f *FileInfoExtended) GetSizeInKB() float64 {
	if f.FsFileInfo == nil {
		f.GetFileInfo()
		return f.SizeBytes / 1024
	}

	f.SizeBytes = float64(f.FsFileInfo.Size())

	return f.SizeBytes / 1024
}

func (f *FileInfoExtended) GetSizeInMB() float64 {
	if f.FsFileInfo == nil {
		f.GetFileInfo()
		return f.SizeBytes / 1048576
	}

	f.SizeBytes = float64(f.FsFileInfo.Size())

	return f.SizeBytes / 1048576
}

func (f *FileInfoExtended) GetSizeInGB() float64 {
	if f.FsFileInfo == nil {
		f.GetFileInfo()
		return f.SizeBytes / 1048576
	}

	f.SizeBytes = float64(f.FsFileInfo.Size())

	return f.SizeBytes / 1073741824
}

func (f *FileInfoExtended) GetFileInfo() error {
	fileinfo, err := os.Stat(f.path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f.FileExists = false
			f.IsDirectory = false
			f.SizeBytes = float64(0.0)
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
	f.SizeBytes = float64(f.FsFileInfo.Size())

	return nil
}

func (f *FileInfoExtended) GetSizeBytes() float64 {
	if f.FsFileInfo == nil {
		f.GetFileInfo()
		return f.SizeBytes
	}

	return f.SizeBytes
}

func (f *FileInfoExtended) CheckSizeBytes() float64 {
	f.GetFileInfo()

	return f.SizeBytes
}
