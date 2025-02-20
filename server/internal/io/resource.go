package io

import (
	"fmt"

	"github.com/GooruApp/gooru/server/internal/config"
	"github.com/c2fo/vfs/vfssimple"
)

func ReadLocalFile(dest []byte, path string) error {
	return readFile(dest, fmt.Sprintf("file://%v", path))
}

func readFile(dest []byte, src string) error {
	osFile, err := vfssimple.NewFile(src)
	if err != nil {
		return err
	}

	if exists, err := osFile.Exists(); err != nil {
		return err
	} else if !exists {
		return fmt.Errorf("path does not exist: %v", src)
	}

	maxSize := config.Settings.MaxFileSize.Get()
	if size, err := osFile.Size(); err != nil {
		return err
	} else if (size / 1024) > uint64(maxSize) {
		return fmt.Errorf("file size (%v kb) exceeds maximum (%v kb)", size/1024, maxSize)
	}

	if _, err := osFile.Read(dest); err != nil {
		return err
	}

	return nil
}
