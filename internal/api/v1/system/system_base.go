package system

import (
	"github.com/gin-gonic/gin"
	systemReq "github.com/rulessly/gin-base/internal/request/system"
	"github.com/rulessly/gin-base/internal/response"
	"github.com/rulessly/gin-base/pkg/utils/captcha"
	"github.com/spf13/cast"
	"slices"
	"time"
)

type BaseApi struct {
}

// Login 账号密码登录
func (b *BaseApi) Login(ctx *gin.Context) {

}

// Register 注册
func (b *BaseApi) Register(ctx *gin.Context) {
	var req systemReq.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数异常,请重新提交", ctx)
		return
	}
	roleList := []int{0x000, 0x011, 0x111}
	if !slices.Contains(roleList, req.Role) {
		response.FailWithMessage("参数异常，不存在的角色", ctx)
		return
	}

	response.Ok(ctx)

}

func (b *BaseApi) Logout(ctx *gin.Context) {}

// Captcha 滑块验证码生成
func (b *BaseApi) Captcha(ctx *gin.Context) {
	captData, err := captcha.SlideCapt.Generate()
	if err != nil {
		response.FailWithMessage("验证码生成失败", ctx)
		return
	}
	dotData := captData.GetData()
	if dotData == nil {
		response.FailWithMessage("验证码生成失败", ctx)
		return
	}

	mBase64, err := captData.GetMasterImage().ToBase64()
	if err != nil {
		response.FailWithMessage("验证码生成失败", ctx)
		return
	}
	tBase64, err := captData.GetTileImage().ToBase64()
	if err != nil {
		response.FailWithMessage("验证码生成失败", ctx)
		return
	}

	resp := map[string]any{
		"timestamp":   cast.ToString(time.Now().Unix()),
		"imageBase64": mBase64,
		"titleBase64": tBase64,
		"titleHeight": dotData.Height,
		"titleWidth":  dotData.Width,
		"titleX":      dotData.DX,
		"titleY":      dotData.DY,
	}

	response.OkWithData(resp, ctx)
}
