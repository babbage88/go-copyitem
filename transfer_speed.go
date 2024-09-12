package main

import (
	"fmt"
	"time"
)

type TransferSpeedLogEntry struct {
	CurrentSizeDest  int       `json:"currentDestinationSize"`
	ChunkSize        int       `json:"chunkSize"`
	ChunkTimeElapsed time.Time `json:"chunkTimeElapsed"`
}

func (f *FileCopyJob) TransferSpeedKB() float64 {
	return f.TransferSpeed / 1024
}

func (f *FileCopyJob) TransferSpeedMB() float64 {
	return f.TransferSpeed / 1048576
}

func (f *FileCopyJob) TransferSpeedGB() float64 {
	return f.TransferSpeed / 1073741824
}

func (f *FileCopyJob) PrettyPrintSpeedBytes() string {
	return fmt.Sprintf("%.2f Bytes/s", f.TransferSpeed)
}

func (f *FileCopyJob) PrettyPrintSpeedKB() string {
	speedKB := f.TransferSpeed / 1024

	return fmt.Sprintf("%.2f KB/s", speedKB)
}

func (f *FileCopyJob) PrettyPrintSpeedMB() string {
	speedMB := f.TransferSpeed / 1048576

	return fmt.Sprintf("%.2f MB/s", speedMB)
}

func (f *FileCopyJob) PrettyPrintSpeedGB() string {
	speedGB := f.TransferSpeed / 1073741824

	return fmt.Sprintf("%.2f GB/s", speedGB)
}
