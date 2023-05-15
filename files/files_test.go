package files

import (
	"fmt"
	"testing"
)

func TestCopyFile(t *testing.T) {
	src := "/tmp/copy/a.txt"
	dst := "/tmp/copy/b.txt"
	err := CopyFile(src, dst, []string{"aaa", "9889898"})
	fmt.Println(err)
}

func TestCopyDir(t *testing.T) {
	src := "/tmp/copy"
	dst := "/tmp/copy1"
	err := CopyDir(src, dst, []string{"aaa", "9889898"}, []string{})
	fmt.Println(err)
}
