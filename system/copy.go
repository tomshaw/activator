package system

import (
	"fmt"
	"github.com/tomshaw/activator/utils"
	"io"
	"io/fs"
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
					return fmt.Errorf("Copy failed %w\n", err)
				}
			}
		}
	}
	return nil
}

func CopyFiles(dst string, files []string) error {
	var err error

	if err = os.MkdirAll(dst, 0664); err != nil {
		return err
	}

	for _, srcfp := range files {
		fileStat, err := os.Stat(srcfp)
		if err != nil {
			return err
		}

		dstfp := path.Join(dst, fileStat.Name())

		if _, ok := utils.SystemFontTypes[filepath.Ext(fileStat.Name())]; ok {
			if err = Copy(srcfp, dstfp); err != nil {
				return fmt.Errorf("Copy failed %w\n", err)
			}
		}
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
