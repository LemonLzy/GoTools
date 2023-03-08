package pretty

type JsonInfo struct {
	json string // 未格式化的json字符串
}

func NewJsonInfo() JsonInfo {
	return JsonInfo{}
}

var jsonInfo = NewJsonInfo()

const module = "[flora][pretty]"
