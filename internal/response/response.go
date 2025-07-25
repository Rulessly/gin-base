package response

import (
	"github.com/gin-gonic/gin"
	"github.com/rulessly/gin-base/pkg/utils/valid"
	"net/http"
)

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func result[T any](code int, msg string, data T, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response[T]{code, msg, data})
}

func Ok(ctx *gin.Context) {
	result(http.StatusOK, "success", map[string]any{}, ctx)
}

func OkWithMessage(msg string, ctx *gin.Context) {
	result(http.StatusOK, msg, map[string]any{}, ctx)
}

func OkWithData[T any](data T, ctx *gin.Context) {
	result(http.StatusOK, "success", data, ctx)
}

func OkWithDetail[T any](msg string, data T, ctx *gin.Context) {
	result(http.StatusOK, msg, data, ctx)
}

func Fail(ctx *gin.Context) {
	result(http.StatusOK, "fail", map[string]any{}, ctx)
}

func FailWithMessage(msg string, ctx *gin.Context) {
	result(http.StatusOK, msg, map[string]any{}, ctx)
}

func FailWithError[T any](err error, obj *T, c *gin.Context) {
	msg := valid.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}
