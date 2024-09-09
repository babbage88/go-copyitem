package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Result struct {
	Error error `json:"error"`
}

type FileCopyJob struct {
	SourceFile        FileInfoExtended `json:"sourcefileinfo"`
	DestinationFile   FileInfoExtended `json:"sourcefileinfo"`
	Running           bool             `json:"jobRunning"`
	Completed         bool             `json:"completed"`
	TimesStarted      int64            `json:"timesStarted"`
	ErrorStatus       error            `json:"status"`
	BytesWritten      int64            `json:"bytesWritten"`
	ProgressCompleted float64          `json:"progress"`
}

type IFileCopyJob interface {
	GetCopyProgressPercentStr() string
	GetCopyProgressPercent() float64
	PrettyPrintSrc() string
	PrettyPrintDst() string
	CopyFile() error
	Start() error
	VerifyDstHash() error
	UpdateProgressBar()
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
	progress := float64(f.BytesWritten) / float64(f.SourceFile.SizeBytes) * 100

	progressStr := fmt.Sprintf("Current Progress: %.2f%%\n", progress)

	return progressStr

}

func (f *FileCopyJob) GetCopyProgressPercentInt64() float64 {
	progress := float64(f.BytesWritten) / float64(f.SourceFile.SizeBytes) * 100

	return progress

}

func (f *FileCopyJob) CopyFile() error {
	fmt.Printf("Starting File Copy Job src: %s\ndst: %s\nsize_kb: %s\n", f.SourceFile.path, f.DestinationFile.path, f.SourceFile.PrettyStringSizeKB())
	src, err := os.Open(f.SourceFile.path)
	if err != nil {
		fmt.Printf("Error Opening source file %s\n", f.PrettyPrintSrc())
		return err
	}
	defer src.Close()

	newfile, err := os.Create(f.DestinationFile.path)
	if err != nil {
		fmt.Printf("Error Creating new destination file: %s\n", f.PrettyPrintDst())
		return err
	}
	defer func() {
		if cerr := newfile.Close(); err == nil {
			err = cerr
		}
	}()

	// Copy the file in chunks and update progress
	buf := make([]byte, 1024*1024) // 1MB buffer
	var totalBytesCopied int64
	srcSize := f.SourceFile.GetSizeBytes()

	for {
		n, readErr := src.Read(buf)
		if n > 0 {
			written, writeErr := newfile.Write(buf[:n])
			if writeErr != nil {
				fmt.Printf("Error writing to destination file: %s\n", f.PrettyPrintDst())
				return writeErr
			}
			totalBytesCopied += int64(written)

			f.BytesWritten = totalBytesCopied
			// Periodically update progress
			progress := float64(totalBytesCopied) / float64(srcSize) * 100
			f.ProgressCompleted = progress

		}

		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			fmt.Printf("Error copying file: %v\n", readErr)
			return readErr
		}
	}

	fmt.Printf("File copy completed: %s -> %s\n", f.PrettyPrintSrc(), f.PrettyPrintDst())
	return nil
}

func (f *FileCopyJob) Start() error {
	f.TimesStarted += 1
	f.Running = true
	fmt.Printf("Starting Copy Job\n")
	fmt.Printf("Source File: %s\n", f.PrettyPrintSrc())
	fmt.Printf("Destination File: %s\n", f.PrettyPrintDst())

	wg := new(sync.WaitGroup)
	wg.Add(2)

	errCh := make(chan error, 1)            // Buffered channel to prevent goroutine block
	boolCompletedChan := make(chan bool, 1) // Buffered channel for completion signal

	// First goroutine for copying the file
	go func() {
		defer wg.Done()

		err := f.CopyFile()
		if err != nil {
			errCh <- err // Send error to errCh if the copy fails
		} else {
			f.Completed = true
			f.Running = false
			boolCompletedChan <- true // Notify of successful completion
		}
		close(boolCompletedChan)
		close(errCh)
	}()

	// Second goroutine for status updates
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-boolCompletedChan:
				return
			case <-ticker.C:
				//fmt.Printf("%s\n", f.GetCopyProgressPercentStr())
				f.UpdateProgressBar()
			}
		}
	}()

	// Main thread, for loop to monitor channels for progress and errors
	for {
		select {
		case <-boolCompletedChan:
			// File copy completed successfully
			fmt.Println("File copy completed.")
			err := f.VerifyDstHash()
			wg.Wait()
			return err
		case err := <-errCh:
			// Return the error if one occurs
			fmt.Printf("Error encountered: %v\n", err)
			wg.Wait() // Ensure all goroutines are done
			return err
		case <-time.After(15 * time.Second):
			// Periodically print which files are being copied.
			//fmt.Printf("Test src: %s dst: %s\n", f.PrettyPrintSrc(), f.PrettyPrintDst())
		}
	}
}

func (f *FileCopyJob) VerifyDstHash() error {
	fmt.Printf("Verifying Destination file Hash matches Source File")
	srcHash, err := f.SourceFile.CalculateFileHash()
	if err != nil {
		return err
	}

	dstHash, err := f.DestinationFile.CalculateFileHash()
	if err != nil {
		return err
	}

	if dstHash != srcHash {
		return fmt.Errorf("file hash mismatch: source and destination files are different")
	} else {
		fmt.Printf("File Hashes Match\n")
		fmt.Printf("Source File: %s\n", f.PrettyPrintSrc())
		fmt.Printf("Source Hash: %s\n", f.SourceFile.FileHash)
		fmt.Printf("Destination File: %s\n", f.PrettyPrintDst())
		fmt.Printf("Destination Hash: %s\n", f.DestinationFile.FileHash)

		return nil
	}

}

func (f *FileCopyJob) UpdateProgressBar() {
	DrawProgressBar(f.ProgressCompleted, 50)
}
