package io

import (
	"bufio"
	"fmt"

	"github.com/GooruApp/gooru/server/internal/config"
	"github.com/c2fo/vfs"
	"github.com/c2fo/vfs/vfssimple"
)

func ReadLocalFile(dest []byte, path string) error {
	return readFile(dest, fmt.Sprintf("file://%v", path))
}

func fileAsReader(file vfs.File) (*bufio.Reader, error) {
	if exists, err := file.Exists(); err != nil {
		return nil, err
	} else if !exists {
		return nil, fmt.Errorf("path does not exist: %v", src)
	}

	maxSize := config.Settings.MaxFileSize.Get()
	if size, err := file.Size(); err != nil {
		return nil, err
	} else if (size / 1024) > uint64(maxSize) {
		return nil, fmt.Errorf("file size (%v kb) exceeds maximum (%v kb)", size/1024, maxSize)
	}

	return bufio.NewReaderSize(file, config.Settings.BufferSize.Get()*1024), nil
}

func readFile(src string) error {
	file, err := vfssimple.NewFile(src)

	if err != nil {
		return err
	}

	defer file.Close()

	bufReader, err := fileAsReader(file)

	if err != nil {
		return err
	}

	for {
		var bytes []byte
		if _, err := bufReader.Read(bytes); err != nil {
			return err
		}
	}
}
