package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
)

type FileInfoExtended struct {
	FsFileInfo  fs.FileInfo `json:"fsfileinfo"`
	path        string      `json:"path"`
	IsDirectory bool        `json:"isDirectory"`
	FileExists  bool        `json:"files_exists"`
	SizeBytes   float64     `json:"sizeBytes"`
	FileHash    string      `json:"fileHash"`
}

type CopyJobFile interface {
	GetFileInfo() error
	GetSizeInKB() float64
	GetSizeInMB() float64
	GetSizeInGB() float64
	GetSizeBytes() float64
	CheckSize() float64
	PrettyStringSizeBytes() string
	PrettyStringSizeKB() string
	PrettyStringSizeMB() string
	PrettyStringSizeGB() string
	CalculateFileHash() (string, error)
	GetFileHash() string
}

func (f *FileInfoExtended) PrettyStringSizeBytes() string {
	coloredsource := fmt.Sprintf("\x1b[%dm%.2f\x1b[0m Bytes", 96, f.GetSizeBytes())
	return coloredsource
}

func (f *FileInfoExtended) PrettyStringSizeKB() string {
	coloredsource := fmt.Sprintf("\x1b[%dm%.2f\x1b[0m KB", 96, f.GetSizeInKB())
	return coloredsource
}

func (f *FileInfoExtended) PrettyStringSizeMB() string {
	coloredsource := fmt.Sprintf("\x1b[%dm%.2f\x1b[0m MB", 96, f.GetSizeInMB())
	return coloredsource
}

func (f *FileInfoExtended) PrettyStringSizeGB() string {
	coloredsource := fmt.Sprintf("\x1b[%dm%.2f\x1b[0m GB", 96, f.GetSizeInGB())
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
		return f.SizeBytes / 1073741824
	}

	f.SizeBytes = float64(f.FsFileInfo.Size())

	return f.SizeBytes / 1073741824
}

func (f *FileInfoExtended) GetFileInfo() error {
	fileinfo, err := os.Stat(f.path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Debug("File does not exist", slog.String("File", f.path))
			f.FileExists = false
			f.IsDirectory = false
			f.SizeBytes = float64(0.0)
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

func (f *FileInfoExtended) CalculateFileHash() (string, error) {
	file, err := os.Open(f.path)
	if err != nil {
		return "", fmt.Errorf("error opening file for hashing: %v", err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("error calculating file hash: %v", err)
	}

	f.FileHash, err = fmt.Sprintf("%x", hash.Sum(nil)), nil

	return f.FileHash, err
}

func (f *FileInfoExtended) GetFileHash() string {
	if f.FileHash == "" {
		f.CalculateFileHash()
	}

	return f.FileHash
}
