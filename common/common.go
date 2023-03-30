package common

import (
	"errors"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

func CheckArgs(cli *cli.Context, command string) error {
	args := cli.Args()
	if args.Len() < 0 {
		color.Red.Printf("%s Command Line new execution error, please use 'flora %s -h for details\n'", module, command)
		return errors.New("lack arguments")
	}
	return nil
}

// IsFileExist 判断文件是否存在
func IsFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// GetFileType 判断文件是什么格式，并返回格式
func GetFileType(path string) string {
	// 实现逻辑
	if path == "FLAC" {

	}
	return "FLAC"
}

// IsAudio 判断是否是音频文件
func IsAudio(contentType string) bool {
	return contentType[:6] == "audio/"
}

// GetAudioFormat 获取音频格式
func GetAudioFormat(path string) string {
	ext := filepath.Ext(path)
	return mime.TypeByExtension(ext)
}

// GetFileContentType 获取文件的 MIME 类型
func GetFileContentType(f *os.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
