package utils

import (
	"os"
)

func ReadFile(filename string) (file []byte, err error) {
	if file, err = os.ReadFile(filename); err != nil {
		return nil, err
	}
	return file, nil
}

func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		return os.IsNotExist(err)
	}
	return true
}

func AppendExists(files []string) []string {
	var fonts []string
	for _, v := range files {
		if FileExists(v) {
			fonts = append(fonts, v)
		}
	}
	return fonts
}
