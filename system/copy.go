package system

import (
	"fmt"
	"github.com/tomshaw/activator/utils"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func CopyFiles(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyFiles(srcfp, dstfp); err != nil {
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

func Copy(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
