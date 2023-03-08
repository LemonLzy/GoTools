package pretty

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"github.com/lemonlzy/flora/common"
	"github.com/urfave/cli/v2"
)

func Action(cli *cli.Context) error {
	err := common.CheckArgs(cli, "pretty")
	if err != nil {
		return err
	}

	json := cli.Args().Get(0)
	fmt.Println(json)
	jsonInfo.json = json
	if json == "" {
		color.Red.Println(module, "json str must match be filled, please check...")
		return errors.New("json str must match be filled")

	}

	formatJson, err := Json(jsonInfo.json)
	if err != nil {
		color.Red.Println(module, "json failed, please check...")
		return errors.New("json failed")
	}

	color.Green.Println(module, "json Success!")
	fmt.Println(formatJson)

	return nil
}
