package service

import (
	"beau-blog/global"
	"beau-blog/model"
	"beau-blog/model/response"
	"go.uber.org/zap"
)

func QueryTagById(id uint)(resTag response.ResTag, err error) {
	if err = global.BB_DB.Model(&model.Tag{}).Select("id", "name", "display_name", "seo_desc", "num").First(&resTag, id).Error; err != nil {
		global.BB_LOG.Error("[GetTagById] error", zap.Any("err: ", err))
		return
	} else {
		return resTag, nil
	}

}

//获取文章的tag list
func QueryTagsByArticleId(articleId uint)(tags []response.ResTag, err error) {
	ids, err := QueryTagListIdsByArticleId(articleId)
	if err != nil {
		return
	}
	for _, id := range ids {
			tag, err := QueryTagById(id)
			if err != nil {
				continue
			}
			tags = append(tags, tag)
	}
	return tags, nil
}
