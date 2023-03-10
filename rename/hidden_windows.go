//go:build windows

package rename

import (
	"os"
	"reflect"
	"strconv"
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
