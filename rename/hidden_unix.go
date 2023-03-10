//go:build darwin || linux

package rename

import (
	"os"
	"strings"
)

func CheckIsHidden(file os.FileInfo) bool {
	filename := file.Name()
	if strings.HasPrefix(filename, ".") {
		return true
	}
	return false
}
