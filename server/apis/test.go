package apis

import (
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 定义 MyController 结构体
type MyController struct {
	Ctx *gin.Context
}
type MyController2 struct {
	Ctx *gin.Context
}

// GetUsers 方法对应于 HTTP 的 GET 请求和 "/mycontroller/users" 路由
func (c *MyController) GetUsersBy(id string) {
	//token, err := utils.CreateToken()
	//if err != nil {
	//	ctx.String(http.StatusOK, "This is a response from GetUsers")
	//	return
	//}
	log.Println(models.User{})
	c.Ctx.String(http.StatusOK, id)
}

// GetUsers 方法对应于 HTTP 的 GET 请求和 "/mycontroller/users" 路由
func (c *MyController2) GetUsers() {
	token, err := utils.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2OTI4NjMwMTEsInVzZXIiOiJKb2huIERvZSJ9.RZSaihLlzFemJXxR_21esXUqMpkwhgQf5gGh2MF4zwk")

	if err != nil {
		c.Ctx.String(http.StatusOK, "This is a response from GetUsers2")
		return
	}
	c.Ctx.JSON(http.StatusOK, token)
}
