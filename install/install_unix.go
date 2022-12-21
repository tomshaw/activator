//go:build unix
// +build unix

package install

import (
	"fmt"
	"github.com/tomshaw/activator/utils"
	"os"
	"path"
)

func install(font *FontData) (err error) {
	target := path.Join(utils.SystemPaths["unix"], path.Base(font.FileName))
	if err = os.MkdirAll(path.Dir(target), 0700); err != nil {
		return fmt.Errorf("Error creating directory: %w", err)
	}
	if err = os.WriteFile(target, font.Data, 0644); err != nil { //nolint
		return fmt.Errorf("Error writing file: %w", err)
	}
	return nil
}

func uninstall(font *FontData) (err error) {
	target := path.Join(utils.SystemPaths["unix"], path.Base(font.FileName))
	if utils.FileExists(target) {
		if err := os.Remove(target); err != nil {
			return fmt.Errorf("Error removing file: %w", err)
		}
	}
	return nil
}

// Windows function stub
func winTempInstall(font *FontData) (err error) {
	return nil
}

// Windows function stub
func winTempUninstall(font *FontData) (err error) {
	return nil
}
