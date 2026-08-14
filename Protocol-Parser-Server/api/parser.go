package api

import (
	"net/http"

	"encoding/hex"
	"strings"

	"github.com/gin-gonic/gin"

	"protocol-parser-server/parser/core"
)

type AnalyzeRequest struct {
	Hex      string `json:"hex"`
	Protocol string `json:"protocol"`
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

	normalized := strings.NewReplacer(" ", "", "\n", "", "\r", "", "\t", "").Replace(req.Hex)
	data, err := hex.DecodeString(normalized)

	if err != nil {

		c.JSON(400, gin.H{

			"error": "HEX格式错误",
		})

		return

	}

	var p core.Protocol
	if req.Protocol != "" {
		p, err = core.Get(req.Protocol)
	} else {
		p, err = core.Detect(data)
	}

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
