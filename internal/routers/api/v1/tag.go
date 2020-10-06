package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zcpin/blog-service/global"
	"github.com/zcpin/blog-service/internal/service"
	"github.com/zcpin/blog-service/pkg/app"
	"github.com/zcpin/blog-service/pkg/errcode"
	"log"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

//@Summary 获取多个标签
//@Produce json
//@Param name query string false "标签名称" maxlength(100)
//@Param state query int false "标签状态"
//@Param page query int false "页码"
//@Param page_size query int false "每页数量"
//@Success 200 {object} model.Tag "成功"
//@Faiure 400 {object} errcode.Error "请求错误"
//@Faiure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	//param := struct {
	//	Name  string `form:"name" binding:"max=100"`
	//	State uint8  `form:"state default=1" binding:"oneof=0 1"`
	//}{}
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	log.Printf("zcp:%+v", param)
	valid, errs := app.BindAndValid(c, &param)
	log.Printf("zcp:%+v", param)
	if valid == true {
		global.Logger.Errorf("app.BindAndValid err %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{})
	return
}

//@Summary 新增标签
//@Produce json
//@Param name body string true "标签名称" minlength(3) maxlength(100)
//@Param state body int false "标签状态" Enums(0,1) default(1)
//@Param created_by body string false "创建者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Faiure 400 {object} errcode.Error "请求错误"
//@Faiure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

//@Summary 更新标签
//@Produce json
//@Param id path int true "标签ID"
//@Param name body string false "标签名称" minlength(3) maxlength(100)
//@Param state body int false "标签状态" Enums(0,1) default(1)
//@Param modified_by body string true "修改者" minlength(3) maxlength(100)
//@Success 200 {object} model.Tag "成功"
//@Faiure 400 {object} errcode.Error "请求错误"
//@Faiure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags/{id} [post]
func (t Tag) Update(c *gin.Context) {}

//@Summary 删除标签
//@Produce json
//@Param id path int true "标签ID"
//@Success 200 {object} model.Tag "成功"
//@Faiure 400 {object} errcode.Error "请求错误"
//@Faiure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
