package tools

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func CheckIsHidden(file os.FileInfo) bool {
	//"通过反射来获取Win32FileAttributeData的FileAttributes
	fa := reflect.ValueOf(file.Sys()).Elem().FieldByName("FileAttributes").Uint()
	byteFa := []byte(strconv.FormatUint(fa, 2))
	if byteFa[len(byteFa)-2] == '1' {
		return true
	}
	return false
}

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
		if strings.HasPrefix(filename, ".") {
			return filepath.SkipDir
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
