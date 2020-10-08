package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zcpin/blog-service/global"
	"github.com/zcpin/blog-service/internal/service"
	"github.com/zcpin/blog-service/pkg/app"
	"github.com/zcpin/blog-service/pkg/convert"
	"github.com/zcpin/blog-service/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

//@Summary 根据指定ID获取文章
//@Produce json
//@Param id path int true "文章标签ID"
//@Param state query int false "文章状态"
//@Success 200 {object} model.Article "成功"
//@Faiure 400 {object} errcode.Error "请求错误"
//@Faiure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}
	response.ToResponse(article)
	return
}

//@Summary 获取文章列表
//@Produce json
//@Param tag_id query int true "文章标签ID"
//@Param state query int false "文章状态"
//@Param page query int false "页码"
//@Param page_size query int false "每页数量"
//@Success 200 {object} model.Article "成功"
//@Faiure 400 {object} errcode.Error "请求错误"
//@Faiure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articles, totalrows, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err: %v", errs)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}
	response.ToResponseList(articles, totalrows)
}

//@Summary 添加文章
//@Produce json
//@Param title body string true "文章标题" minlength(4) maxlength(100)
//@Param desc body string true "文章简介" minlength(10) maxlength(255)
//@Param content body string true "文章内容" minlength(10)
//@Param cover_image_url body string true "封面图片链接" maxlength(255)
//@Param state body int false "文章状态" Enums(0,1) default(1)
//@Param created_by body string true "创建者" minlength(2)
//@Param tag_id body string true "文章标签ID"
//@Success 200 {object} model.Article "成功"
//@Faiure 400 {object} errcode.Error "请求错误"
//@Faiure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err: %v", errs)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

//@Summary 更新文章
//@Produce json
//@Param id path int true "文章ID"
//@Param title body string false "文章标题" minlength(4) maxlength(100)
//@Param desc body string false "文章简介" minlength(10) maxlength(255)
//@Param content body string false "文章内容" minlength(10)
//@Param cover_image_url body string false "封面图片链接" maxlength(255)
//@Param state body int false "文章状态" Enums(0,1) default(1)
//@Param tag_id body string true "文章标签ID"
//@Param modified_by body string true "修改者" minlength(3) maxlength(100)
//@Success 200 {object} model.Article "成功"
//@Faiure 400 {object} errcode.Error "请求错误"
//@Faiure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/article/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateActicleFial)
		return
	}
	response.ToResponse(gin.H{})
	return
}

//@Summary 删除文章
//@Produce json
//@Param id path int true "文章ID"
//@Success 200 {object} model.Article "成功"
//@Faiure 400 {object} errcode.Error "请求错误"
//@Faiure 500 {object} errcode.Error "内部错误"
//@Router /api/v1/article/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}

	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	if valid == true {
		global.Logger.Errorf("app.BindAndValid err: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err:% v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
