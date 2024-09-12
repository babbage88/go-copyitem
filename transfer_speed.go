package main

import (
	"fmt"
)

func (f *FileCopyJob) PrettyPrintSpeedBytes() string {
	return fmt.Sprintf("Speed: %.2f Bytes/s", f.TransferSpeed)
}

func (f *FileCopyJob) PrettyPrintSpeedKB() string {
	speedKB := f.TransferSpeed / 1024

	return fmt.Sprintf("Speed: %.2f KB/s", speedKB)
}

func (f *FileCopyJob) PrettyPrintSpeedMB() string {
	speedMB := f.TransferSpeed / 1048576

	return fmt.Sprintf("Speed: %.2f MB/s", speedMB)
}

func (f *FileCopyJob) PrettyPrintSpeedGB() string {
	speedGB := f.TransferSpeed / 1073741824

	return fmt.Sprintf("Speed: %.2f GB/s", speedGB)
}
