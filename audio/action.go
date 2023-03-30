package audio

import (
	"fmt"
	"github.com/lemonlzy/flora/common"
	"github.com/urfave/cli/v2"
)

func Action(cli *cli.Context) error {
	err := common.CheckArgs(cli, "audio")
	if err != nil {
		return err
	}

	path := cli.Args().Get(0)
	fmt.Println("path:", path)

	// 判断当前路径下文件是否存在
	if ok := common.IsFileExist(path); !ok {
		return fmt.Errorf("file not exist")
	}

	// 根据音频格式创建对应的结构体

	return nil
}
