package service

import (
	"beau-blog/global"
	"beau-blog/model"
	"beau-blog/model/request"
	"beau-blog/model/response"
	"go.uber.org/zap"
)

// 获取文章列表
func QueryArticleList(info request.ReqPageInfo, category uint, tag uint, order string, desc bool) (err error, list []model.Article, count int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.BB_DB.Model(&model.Article{})
	err = db.Find(&model.Article{}).Count(&count).Error
	if err != nil {
		return
	} else {
		err = db.Limit(limit).Offset(offset).Find(&list).Error
	}
	return
}

// 获取文章详情
func QueryArticle(id uint) (articleDetail response.ReqArticleDetail, err error) {
	var article model.Article
	if err = global.BB_DB.Model(&model.Article{}).First(&article, id).Error; err != nil {
		return
	} else {
		// get category
		cate, err := QueryCategoryByArticleId(article.ID)
		if err != nil {
			return
		}
		articleDetail.Category = cate
		// get tag
		tags, err := QueryTagsByArticleId(article.ID)
		if err != nil {
			return
		}
		articleDetail.Tags = tags
		return articleDetail, nil
	}
}

// 创建文章
func CreateArticle(reqArticle request.ReqArticle, userId uint) (resArticle response.ReqArticleDetail, err error) {
	var article = model.Article{
		UserId:       userId,
		Title:        reqArticle.Title,
		Content:      reqArticle.Content,
		AllowComment: reqArticle.AllowComment,
		IsPublic:     reqArticle.IsPublic,
	}
	if err = global.BB_DB.Model(&model.Article{}).Create(&article).Error; err != nil {
		global.BB_LOG.Error("[CreateArticle] error", zap.Any("err: ", err))
		return
	} else {
		// 写入article - category
		err = CreateCorrelation(article, reqArticle.Category, model.CorelationArticleCategory)
		if err != nil {
			return 
		}
		// 写入article - tag
		if len(reqArticle.Tags) != 0 {
			for _,v := range reqArticle.Tags {
				err = CreateCorrelation(article, v, model.CorelationArticleTag)
				if err != nil {
					continue
				}
			}
		}

		resArticle, err = QueryArticle(article.ID)
		if err != nil {
			return
		}
		return resArticle, nil
	}
}

// 编辑文章

// 删除文章
