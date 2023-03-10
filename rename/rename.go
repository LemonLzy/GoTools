package rename

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func BatchRenameFiles(dir, subStr, newName string) error {
	var filenameSlice []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err.Error())
			return err
		}

		if info.IsDir() {
			return nil
		}

		filename := info.Name()
		if CheckIsHidden(info) {
			// 以.开头的文件，或者隐藏文件直接跳过
			return nil
		}

		fileDir := filepath.Dir(path) + "/"
		err = CheckIsKeyword(filename, subStr, fileDir, newName)
		if err != nil {
			return err
		}

		filenameSlice = append(filenameSlice, filename)
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func CheckIsKeyword(filename, subStr, dir, newName string) error {
	if strings.Contains(filename, subStr) {
		err := ReplaceFilename(filename, subStr, dir, newName)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReplaceFilename(filename, subStr, dir, newName string) error {
	dir = filepath.ToSlash(dir)
	err := os.Rename(dir+filename, dir+strings.ReplaceAll(filename, subStr, newName))
	if err != nil {
		return err
	}
	return nil
}

func IsPathExist(path string) bool {
	// 必须为绝对路径
	if ok := filepath.IsAbs(path); !ok {
		return false
	}

	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false

	}
	return true
}

// IsDir 判断是否是文件夹
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return stat.IsDir()
}

// HasReadPermission 判断文件夹是否具有可读权限
func HasReadPermission(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		color.Red.Println(module, "get permission failed, please check...")
		return false
	}

	fileMode := stat.Mode()
	fmt.Println("file_mode:", fileMode)
	perm := fileMode.Perm()
	fmt.Println("permission:", uint32(perm))

	return false
}

// GrantReadPermission 赋予文件夹可读权限
func GrantReadPermission(path string) error {
	err := os.Chmod(path, 0744)
	if err != nil {
		color.Red.Println(module, "grant permission failed, please check...")
		return errors.New("grant Permission failed")
	}
	return nil
}
