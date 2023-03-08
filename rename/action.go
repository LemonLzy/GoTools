package rename

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"github.com/lemonlzy/flora/common"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

func Action(cli *cli.Context) error {
	err := common.CheckArgs(cli, "rename")
	if err != nil {
		return err
	}

	path := cli.Args().Get(0)
	fmt.Println("path:", path)

	// 路径必须为绝对路径
	if ok := filepath.IsAbs(path); !ok {
		color.Red.Println(module, "path must be absolute path")
		return errors.New("path must be absolute path")
	}

	// 必须为文件夹，不能为文件路径
	if ok := IsDir(path); !ok {
		color.Red.Println(module, "path must be Dir, please check...")
		return errors.New("path must be Dir")
	}

	removeName := cli.Args().Get(1)
	fmt.Println("removeName:", removeName)

	// 需要移除的名称不能为空
	if removeName == "" {
		color.Red.Println(module, "remove_name must match be filled, please check...")
		return errors.New("remove_name must match be filled")
	}

	newName := cli.Args().Get(2)

	// 批量重命名文件夹下，带有removeName字符串的文件
	err = BatchRenameFiles(path, removeName, newName)
	if err != nil {
		color.Red.Println(module, "BatchRename failed, please check...")
		return errors.New("BatchRename failed")
	}

	color.Green.Println(module, "BatchRename Success!")

	return nil
}
