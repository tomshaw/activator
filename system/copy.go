package system

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/tomshaw/activator/utils"
)

func CopyFiles(source, destination string) error {
	err := filepath.WalkDir(source, func(file string, item fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Unsupported mime type: %w", err)
		}
		if _, ok := utils.SystemFontTypes[filepath.Ext(item.Name())]; ok {
			dst := path.Join(destination, item.Name())
			_, err := copy(file, dst)
			if err != nil {
				return fmt.Errorf("Copy failed %w\n", err)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("WalkDir process error: %w", err)
	}
	return nil
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)

	return nBytes, err
}
