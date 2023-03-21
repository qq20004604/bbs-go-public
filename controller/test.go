package controller

import (
	"github.com/gin-gonic/gin"
	"main/utils"
)

type TestForm struct {
	ID uint `json:"id" label:"id"`
}

func Test(c *gin.Context) {
	utils.SuccessJson(c, "test", nil)
}
