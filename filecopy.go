package main

import (
	"crypto/sha256"
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
	SourceFile      FileInfoExtended `json:"sourcefileinfo"`
	DestinationFile FileInfoExtended `json:"sourcefileinfo"`
	Running         bool             `json:"jobRunning"`
	Completed       bool             `json:"completed"`
	TimesStarted    int64            `json:"timesStarted"`
	ErrorStatus     error            `json:"status"`
}

type IFileCopyJob interface {
	GetCopyProgressPercentStr() string
	GetCopyProgressPercentInt64() int64
	PrettyPrintSrc() string
	PrettyPrintDst() string
	CopyFile() error
	Start() error
	VerifyDstHash() error
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
	dstlen := int64(f.DestinationFile.CheckSizeBytes())
	srclength := int64(f.SourceFile.CheckSizeBytes())

	cursizedec := dstlen / srclength
	cursize := cursizedec * 100
	cursizestr := fmt.Sprint(cursize, "%")

	return cursizestr

}

func (f *FileCopyJob) GetCopyProgressPercentInt64() int64 {
	dstlen := int64(f.DestinationFile.CheckSizeBytes())
	srclength := int64(f.SourceFile.CheckSizeBytes())

	cursizedec := dstlen / srclength
	cursize := cursizedec * 100

	return cursize

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

			// Periodically update progress
			progress := float64(totalBytesCopied) / float64(srcSize) * 100
			fmt.Printf("Current Progress: %.2f%%\n", progress)
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
		time.Sleep(8 * time.Second)
		fmt.Printf("Test src: %s dst: %s\n", f.PrettyPrintSrc(), f.PrettyPrintDst())
	}()

	// Main thread, for loop to monitor channels for progress and errors
	for {
		select {
		case <-boolCompletedChan:
			// File copy completed successfully
			fmt.Println("File copy completed.")
			wg.Wait()
			return nil
		case err := <-errCh:
			// Return the error if one occurs
			fmt.Printf("Error encountered: %v\n", err)
			wg.Wait() // Ensure all goroutines are done
			return err
		case <-time.After(5 * time.Second):
			// Periodically print progress
			fmt.Printf("Current Progress: %s\n", f.GetCopyProgressPercentStr())
		}
	}
}

func (f *FileCopyJob) VerifyDstHash() error {
	src, err := os.Open(f.SourceFile.path)

	if err != nil {
		fmt.Errorf("Error Operning %s", src)
	}
	defer src.Close()

	dst, err := os.Open(f.DestinationFile.path)

	if err != nil {
		fmt.Errorf("Error Operning %s", dst)
	}
	defer src.Close()

	srcHash := sha256.New()
	if _, err := io.Copy(srcHash, src); err != nil {
		fmt.Errorf("Error Hashing file %s", src)
	}

	dstHash := sha256.New()
	if _, err := io.Copy(dstHash, src); err != nil {
		fmt.Errorf("Error Hashing file %s", dst)
	}

	if srcHash != dstHash {
		retVal := fmt.Errorf("Hashes dot match src: %s dst: %s", src, dst)
		return retVal
	} else {
		fmt.Printf("Destination File hash matches source\n")
		return nil
	}

}
