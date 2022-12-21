package install

import (
	"bytes"
	"fmt"
	"github.com/ConradIrwin/font/sfnt"
	"github.com/tomshaw/activator/utils"
	"path"
)

type FontData struct {
	FileName string
	FilePath string
	FontName string
	MimeType string
	Data     []byte
}

func Read(filePath string) (fontData *FontData, err error) {
	var byteData []byte

	byteData, err = utils.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileName := path.Base(filePath)
	if _, ok := utils.InstallFontTypes[path.Ext(fileName)]; !ok {
		return nil, fmt.Errorf("Unsupported mime type: %v", fileName)
	}

	regString := ""
	if path.Ext(fileName) == ".otf" {
		regString = "(OpenType)"
	} else {
		regString = "(TrueType)"
	}

	font, err := sfnt.Parse(bytes.NewReader(byteData))
	if err != nil {
		return nil, fmt.Errorf("Error reading font %v", fileName)
	}

	if !font.HasTable(sfnt.TagName) {
		return nil, fmt.Errorf("Error reading name table %v", fileName)
	}

	items, err := font.NameTable()
	if err != nil {
		return nil, fmt.Errorf("Error reading name table %v", fileName)
	}

	records := make(map[sfnt.NameID]string)
	for _, item := range items.List() {
		records[item.NameID] = item.String()
	}

	fontName := records[sfnt.NameFull]
	if len(fontName) == 0 {
		return nil, fmt.Errorf("Error reading font name %v", fileName)
	}

	fontData = &FontData{
		FileName: fileName,
		FilePath: filePath,
		FontName: fontName,
		MimeType: regString,
		Data:     byteData,
	}

	return fontData, nil
}
