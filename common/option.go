package common

type ProjectInfo struct {
	Name       string // 项目名
	ParentPath string // 项目父目录
	RootPath   string // parent_path + name，项目根目录，绝对路径
	ModuleName string // 项目模块名，go.mod文件中的module
}

func NewProjectInfo() ProjectInfo {
	return ProjectInfo{}
}

const module = "[flora][COMMON]"
