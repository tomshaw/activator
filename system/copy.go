package system

import (
	"fmt"
	"github.com/tomshaw/activator/utils"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

func CopyFilesFolders(src string, dst string) error {
	var err error
	var fds []fs.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyFilesFolders(srcfp, dstfp); err != nil {
				return fmt.Errorf("Copy failed %w\n", err)
			}
		} else {
			if _, ok := utils.SystemFontTypes[filepath.Ext(fd.Name())]; ok {
				if err = Copy(srcfp, dstfp); err != nil {
					log.Fatalf("Copy failed %v", err)
				}
			}
		}
	}
	return nil
}

func CopyFiles(src, dst string) error {
	err := filepath.WalkDir(src, func(src string, item fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Unsupported mime type: %w", err)
		}
		if _, ok := utils.SystemFontTypes[filepath.Ext(item.Name())]; ok {
			dst := path.Join(dst, item.Name())
			err := Copy(src, dst)
			if err != nil {
				log.Fatalf("Copy failed %v", err)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("WalkDir process error: %w", err)
	}
	return nil
}

func Copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	srcfd, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcfd.Close()

	dstfd, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}

	return nil
}
