package service

import (
	"beau-blog/global"
	"beau-blog/model"
	"go.uber.org/zap"
)

func QueryCategoryIdByArticleId(articleId uint) (cateId uint, err error) {
	var corellation model.Corelation
	if err = global.BB_DB.Model(&model.Corelation{}).Where("id_1 = ? && type = ?", articleId, model.CorelationArticleCategory).First(&corellation).Error; err != nil {
		global.BB_LOG.Error("[GetCategoryIdByArticleId] error", zap.Any("err: ", err))
		return
	} else {
		return corellation.ID2,nil
	}
}

func QueryTagListIdsByArticleId(articleId uint)(tags []uint, err error) {
	var corellations []model.Corelation
	if err = global.BB_DB.Model(&model.Corelation{}).Where("id_1 = ? && type = ?", articleId, model.CorelationArticleTag).Find(&corellations).Error; err != nil {
		global.BB_LOG.Error("[GetTagListIdsByArticleId] error", zap.Any("err: ", err))
		return
	} else {
		for _,v := range corellations {
			tags = append(tags, v.ID2)
		}
		return tags, nil
	}
}

func CreateCorrelation(articleId uint, id2 uint, rType int) (err error) {
	var correlation = model.Corelation{
		ID1: articleId,
		ID2: id2,
		Type:rType,
	}
	if err = global.BB_DB.Model(&model.Corelation{}).Create(&correlation).Error; err != nil {
		global.BB_LOG.Error("[CreateCorrelation] error", zap.Any("err: ", err))
		return
	}
	return
}