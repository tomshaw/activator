package utils

import (
	"os"
	"path"
)

var SystemPaths = map[string]string{
	"unix":    path.Join(os.Getenv("HOME"), "/.local/share/fonts"),
	"windows": path.Join(os.Getenv("WINDIR"), "Fonts"),
	"darwin":  path.Join(os.Getenv("HOME"), "/Library/Fonts"),
}
