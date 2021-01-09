package service

import (
	"beau-blog/global"
	"beau-blog/model/response"
	"go.uber.org/zap"
)

/**
* 获取子分类
 */
func QueryCategoryByParentId(parentId uint)(cates []response.ResCategory, err error) {
	if err = global.BB_DB.Table("articles").Select("id", "name", "display_name", "seo_desc").Find(&cates).Error; err != nil {
		global.BB_LOG.Error("GetCategoryByParentId error", zap.Any("err: ", err) )
		return
	} else {
		return cates, nil
	}
}

/**
* 通过id 获取分类
 */
func QueryCategoryById(id uint)(resCate response.ResCategory, err error) {
	if err = global.BB_DB.Table("articles").Select("id", "name", "display_name", "seo_desc").Scan(&resCate).Error; err != nil {
		global.BB_LOG.Error("[GetCategoryById] error", zap.Any("err: ", err) )
		return
	} else {
		return resCate, nil
	}
}

func QueryCategoryByArticleId(id uint)(resCate response.ResCategory, err error) {
	cateId, err := QueryCategoryIdByArticleId(id)
	if err != nil {
		global.BB_LOG.Error("[GetCategoryByArticleId] error", zap.Any("err: ", err) )
		return
	}
	resCate, err = QueryCategoryById(cateId)
	return
}

