package system

import (
	"fmt"
	"github.com/tomshaw/activator/utils"
	"io/fs"
	"path/filepath"
)

func FindFiles(root string) ([]string, error) {
	var found []string
	err := filepath.WalkDir(root, func(f string, item fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("WalkDir error: %w", err)
		}
		if _, ok := utils.SystemFontTypes[filepath.Ext(item.Name())]; ok {
			found = append(found, f)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return found, nil
}
