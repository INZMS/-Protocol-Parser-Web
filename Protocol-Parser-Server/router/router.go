// 路由初始化
package router

import (
	"github.com/gin-gonic/gin"

	"protocol-parser-server/api"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	api.RegisterParserRouter(r)

	return r

}
