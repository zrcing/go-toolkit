package files

import (
	"bytes"
	"github.com/zrcing/go-toolkit/vector"
	"os"
	"path/filepath"
)

func CopyFile(source string, dest string, replaces []string) error {
	srcInfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	buffer, err := os.ReadFile(source)
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
	return os.WriteFile(dest, buffer, srcInfo.Mode())
}

func CopyDir(source string, dest string, replaces, ignores []string) error {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, sourceInfo.Mode())
	if err != nil {
		return err
	}

	fileDirs, err := os.ReadDir(source)
	if err != nil {
		return err
	}
	for _, fileDir := range fileDirs {
		if vector.InVectors(fileDir.Name(), ignores) {
			continue
		}
		srcFilePath := filepath.Join(source, fileDir.Name())
		dstFilePath := filepath.Join(dest, fileDir.Name())
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
