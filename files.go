package toolkit

import (
	"bytes"
	"os"
	"path/filepath"
)

func CopyFile(src string, dst string, replaces []string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	buffer, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	var oldStr string
	for i, v := range replaces {
		if i%2 == 0 {
			oldStr = v
			continue
		}
		buffer = bytes.ReplaceAll(buffer, []byte(oldStr), []byte(v))
	}
	return os.WriteFile(dst, buffer, srcInfo.Mode())
}

func CopyDir(src string, dst string, replaces, ignores []string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	fileDirs, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, fileDir := range fileDirs {
		if InVectors(fileDir.Name(), ignores) {
			continue
		}
		srcFilePath := filepath.Join(src, fileDir.Name())
		dstFilePath := filepath.Join(dst, fileDir.Name())
		var e error
		if fileDir.IsDir() {
			e = CopyDir(srcFilePath, dstFilePath, replaces, ignores)
		} else {
			e = CopyFile(srcFilePath, dstFilePath, replaces)
		}
		if e != nil {
			return e
		}
	}
	return nil
}
