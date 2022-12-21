package install

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func Init(files []string, activate bool, temp bool) error {
	var errors []string
	for _, v := range files {
		if err := run(filepath.ToSlash(v), activate, temp); err != nil {
			errors = append(errors, fmt.Errorf("Font installation error:%w", err).Error())
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}
	return nil
}

func run(fontPath string, activate bool, temp bool) (err error) {
	var font *FontData

	font, err = Read(fontPath)
	if err != nil {
		return err
	}

	if temp && runtime.GOOS == "windows" {
		if activate {
			err = winTempInstall(font)
			if err == nil {
				return err
			}
		} else {
			err = winTempUninstall(font)
			if err == nil {
				return err
			}
		}
	} else {
		if activate {
			err = install(font)
			if err == nil {
				return err
			}
		} else {
			err = uninstall(font)
			if err == nil {
				return err
			}
		}
	}

	return nil
}
