package main

type SizeInKb interface {
	GetSizeInKB() int64
}

type SizeInMb interface {
	GetSizeInMB() int64
}

func (f *FileInfoExtended) GetSizeInKB() int64 {
	flength := f.FsFileInfo.Size()
	kbsize := flength / 1024

	return kbsize
}

func (f *FileInfoExtended) GetSizeInMB() int64 {
	flength := f.FsFileInfo.Size()
	mbsize := flength / 1048576

	return mbsize
}
