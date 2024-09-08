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
	fmt.Printf("Starting File Copy Job src: %s\ndst: %s\nsize_kb: %s\n", f.SourceFile.path, f.DestinationFile.path, f.SourceFile.GetSizeInKB())
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

func (f *FileCopyJob) Start() error {
	f.TimesStarted += 1
	f.Running = true
	fmt.Printf("Starting Copy Job")
	fmt.Printf("Source File: %s\n", f.PrettyPrintSrc())
	fmt.Printf("Destination File: %s\n", f.PrettyPrintDst())

	wg := new(sync.WaitGroup) //create a wait group
	wg.Add(2)

	errCh := make(chan error)
	boolCompletedChan := make(chan bool)

	// Start Copying the file, return any errors the the Errors Channel
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		err := f.CopyFile()
		if err != nil {
			errCh <- err
		}

		f.Completed = true
		f.Running = false
		boolCompletedChan <- true
		close(boolCompletedChan)

	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(8 * time.Second)
		fmt.Printf("Test src: %s dst: %s\n", f.PrettyPrintSrc(), f.PrettyPrintDst())
	}(wg)

	for !f.Completed {
		time.Sleep(15 * time.Second)
		fmt.Printf("Current Progress %s\n", f.GetCopyProgressPercentStr())
	}

	wg.Wait()
	return nil // temporary return nil to twst
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
