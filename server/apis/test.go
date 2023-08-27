package apis

import (
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 定义 MyController 结构体
type MyController struct {
	Ctx *gin.Context
}
type MyController2 struct {
	Ctx *gin.Context
}

// GetUsersBy 方法对应于 HTTP 的 GET 请求和 "/mycontroller/users" 路由
func (c *MyController) GetUsersBy(id string) {
	// string 转 int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.Ctx.String(http.StatusOK, "This is a response from GetUsers")
		return
	}
	encryptID := utils.EncryptID(idInt)

	c.Ctx.String(http.StatusOK, encryptID)
}

// GetTestBy 方法对应于 HTTP 的 GET 请求和 "/mycontroller/users" 路由
func (c *MyController) GetTestBy(id string) {
	decryptID, err := utils.DecryptID(id)

	if err != nil {
		c.Ctx.String(http.StatusOK, "This is a response from GetUsers")
		return
	}
	c.Ctx.String(http.StatusOK, strconv.Itoa(decryptID))
}

// GetUsers 方法对应于 HTTP 的 GET 请求和 "/mycontroller/users" 路由
func (c *MyController2) GetUsers() {
	token, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTI4NjMwMTEsInVzZXIiOiJKb2huIERvZSJ9.RZSaihLlzFemJXxR_21esXUqMpkwhgQf5gGh2MF4zwk")

	if err != nil {
		c.Ctx.String(http.StatusOK, "This is a response from GetUsers2")
		return
	}
	c.Ctx.JSON(http.StatusOK, token)
}
