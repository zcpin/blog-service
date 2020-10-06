package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zcpin/blog-service/pkg/app"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	data:=map[string]interface{}{
		"code" : 0,
		"msg" : "成功",
	}
	app.NewResponse(c).ToResponse(data)
	//app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
