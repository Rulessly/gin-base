package test

import (
	"encoding/json"
	"fmt"
	"github.com/rulessly/gin-base/pkg/utils/captcha"
	"log"
	"testing"
	"time"
)

// 滑动验证码测试
func TestCaptchaSlide(t *testing.T) {
	fmt.Println(time.Now().Unix())
	captData, err := captcha.SlideCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	dotData := captData.GetData()
	if dotData == nil {
		log.Fatalln(">>>>> generate err")
	}

	dots, _ := json.Marshal(dotData)
	fmt.Println(">>>>> ", string(dots))

	var mBase64, tBase64 string
	mBase64, err = captData.GetMasterImage().ToBase64()
	if err != nil {
		fmt.Println(err)
	}
	tBase64, err = captData.GetTileImage().ToBase64()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(">>>>> ", mBase64)
	fmt.Println(">>>>> ", tBase64)

}
