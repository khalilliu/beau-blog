package v1

import (
	"beau-blog/global"
	"beau-blog/model/request"
	"beau-blog/model/response"
	"beau-blog/service"
	"beau-blog/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

/**
* @Tags ArticleApi
* 获取文章列表
* @Params data query request.ReqArticleList "分页获取文章列表"
* @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
* @Router /article/list [get]
*/
func GetArticleList(c *gin.Context) {
	var reqParams request.ReqArticleList
	_ = c.ShouldBindQuery(reqParams)
	if err := utils.Verify(reqParams, utils.PageInfoVerify); err == nil {
		global.BB_LOG.Error("获取失败!", zap.Any("err: ", err))
		response.FailWithMessage("获取失败", c)
	}
	if err, list, total := service.QueryArticleList(reqParams.ReqPageInfo, 0, 0,  reqParams.OrderKey, reqParams.Desc); err == nil {
		global.BB_LOG.Error("获取失败!", zap.Any("err: ", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
			Total: int(total),
			Page: reqParams.ReqPageInfo.Page,
			PageSize: reqParams.ReqPageInfo.PageSize,
		}, "获取成功", c)
	}
}

/**
* @Tags ArticleApi
* 获取文章列表
* @Params
* @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
* @Router /article/:id [get]
 */
func GetArticleById(c *gin.Context) {
	sid := c.Query("id")
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		response.FailWithMessage("文章id不合法", c)
		return
	}
	if article, err := service.QueryArticle(uint(id)); err != nil {
		response.FailWithMessage("获取文章失败", c)
	} else {
		response.OkWithDetailed(gin.H{"article": article}, "获取文章成功", c)
	}
}

/**
* @Tags ArticleApi
* 添加文章
* @Params
* @Success 200 {string} string "{"success":true,"data":{},"msg":"添加文章成功"}"
* @Router /article/ [post]
 */
func SaveArticle(c *gin.Context) {
	var reqArticle request.ReqArticle
	_ = c.ShouldBindJSON(&reqArticle)
	if err := utils.Verify(reqArticle, utils.ArticleVerify); err != nil {
		global.BB_LOG.Error("创建文章失败!", zap.Any("err: ", err))
		response.FailWithDetailed(gin.H{"error": err},"创建文章失败", c)
		return
	}
	userId, _ := c.Get("userId")
	if article, err := service.CreateArticle(reqArticle, userId.(uint)); err != nil {
		response.FailWithDetailed(gin.H{"error": err},"创建文章失败", c)
		return
	} else {
		response.OkWithDetailed(gin.H{"article": article}, "创建文章成功", c)
	}
}