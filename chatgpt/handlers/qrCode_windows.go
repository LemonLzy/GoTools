//go:build windows

package handlers

import "github.com/eatmoreapple/openwechat"

func QrCodeCallBack(uuid string) {
	openwechat.PrintlnQrcodeUrl(uuid)
}
