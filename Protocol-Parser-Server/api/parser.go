package api

import (
	"net/http"

	"encoding/hex"

	"github.com/gin-gonic/gin"

	"protocol-parser-server/parser/core"
)

type AnalyzeRequest struct {
	Hex string `json:"hex"`
}

// 注册解析路由
func RegisterParserRouter(r *gin.Engine) {

	r.POST("/api/parser/analyze", analyze)

}
func analyze(c *gin.Context) {

	var req AnalyzeRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {

		c.JSON(400, gin.H{

			"error": err.Error(),
		})

		return

	}

	data, err := hex.DecodeString(req.Hex)

	if err != nil {

		c.JSON(400, gin.H{

			"error": "HEX格式错误",
		})

		return

	}

	p, err := core.Detect(data)

	if err != nil {

		c.JSON(400, gin.H{

			"error": err.Error(),
		})

		return

	}

	result, err := p.Parse(data)

	//增加解析错误判断
	if err != nil {

		c.JSON(400, gin.H{
			"success": false,

			"error": err.Error(),
		})

		return

	}
	c.JSON(http.StatusOK, gin.H{

		"success": true,

		"data": result,
	})

}
