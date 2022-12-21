package installer

import (
	"fmt"
	"github.com/tomshaw/activator/utils"
	"golang.org/x/sys/windows/registry"
	"os"
	"path"
	"syscall"
	"unsafe"
)

const (
	HWND_BROADCAST = uintptr(0xFFFF)
	WM_FONTCHANGE  = uintptr(0x001D)
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")
	gdi32  = syscall.NewLazyDLL("gdi32.dll")
)

func broadcastFontChange() error {
	_, _, err := user32.NewProc("SendMessageW").Call(HWND_BROADCAST, WM_FONTCHANGE, 0, 0)
	if err != nil {
		return err
	}
	return nil
}

func addFontResource(filePath string) error {
	strPtr, _ := syscall.UTF16PtrFromString(filePath)
	_, _, err := gdi32.NewProc("AddFontResourceW").Call(uintptr(unsafe.Pointer(strPtr)))
	if err != nil {
		return err
	}
	return nil
}

func removeFontResource(filePath string) error {
	strPtr, _ := syscall.UTF16PtrFromString(filePath)
	_, _, err := gdi32.NewProc("RemoveFontResourceW").Call(uintptr(unsafe.Pointer(strPtr)))
	if err != nil {
		return err
	}
	return nil
}

func install(font *FontData) (err error) {
	target := path.Join(utils.SystemPaths["windows"], font.FileName)
	value := fmt.Sprintf("%v %v", font.FontName, font.MimeType)

	if !utils.FileExists(target) {
		if err = os.WriteFile(target, font.Data, 0o644); err != nil { //nolint
			return fmt.Errorf("Error writing file: %w", err)
		}
	}

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Fonts`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()

	if _, _, err := k.GetStringValue(value); err != nil {
		if err = k.SetStringValue(value, font.FileName); err != nil {
			return fmt.Errorf("Error setting registry value: %w", err)
		}
	}

	if err = addFontResource(font.FilePath); err != nil {
		return fmt.Errorf("Error adding font resource: %w", err)
	}

	if err = broadcastFontChange(); err != nil {
		return fmt.Errorf("Error running broadcast change: %w", err)
	}

	return nil
}

func uninstall(font *FontData) (err error) {
	target := path.Join(utils.SystemPaths["windows"], font.FileName)
	value := fmt.Sprintf("%v %v", font.FontName, font.MimeType)

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Fonts`, registry.ALL_ACCESS)
	if err != nil {
		return fmt.Errorf("Error opeing registry: %w", err)
	}
	defer k.Close()

	if _, _, err = k.GetStringValue(value); err != nil {
		return fmt.Errorf("Error getting registry value: %w", err)
	}

	if err = k.DeleteValue(value); err != nil {
		return fmt.Errorf("Error deleting registry value: %w", err)
	}

	if utils.FileExists(target) {
		if err := os.Remove(target); err != nil {
			return fmt.Errorf("Error removing font file: %w", err)
		}
	}

	if err = removeFontResource(font.FilePath); err != nil {
		return fmt.Errorf("Error removing font resource: %w", err)
	}

	if err = broadcastFontChange(); err != nil {
		return fmt.Errorf("Error running broadcast change: %w", err)
	}

	return nil
}

func WinTempInstall(font *FontData) (err error) {
	if err = addFontResource(font.FilePath); err != nil {
		return fmt.Errorf("Error adding font resource: %w", err)
	}
	if err = broadcastFontChange(); err != nil {
		return fmt.Errorf("Error running broadcast change: %w", err)
	}
	return nil
}

func WinTempUninstall(font *FontData) (err error) {
	if err = removeFontResource(font.FilePath); err != nil {
		return fmt.Errorf("Error adding font resource: %w", err)
	}
	if err = broadcastFontChange(); err != nil {
		return fmt.Errorf("Error running broadcast change: %w", err)
	}
	return nil
}
