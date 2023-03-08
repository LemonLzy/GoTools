package rename

type DirInfo struct {
	Name     string // 目录名
	Path     string // 项目副路径
	RootPath string // path + name，目录绝对路径
}

func NewDirInfo() DirInfo {
	return DirInfo{}
}

var dir = NewDirInfo()

const module = "[flora][tools]"
