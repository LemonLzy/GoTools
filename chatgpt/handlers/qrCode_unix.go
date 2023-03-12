//go:build darwin || linux

package handlers

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
)

func QrCodeCallBack(uuid string) {
	log.Println("login in linux")
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}
