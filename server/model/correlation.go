package model

import "beau-blog/global"

const (
	CorelationArticleCategory = iota
	CorelationArticleTag
	CorelationCateogoryTag
	CorelationArticleArchive
)

// Corelation model.
// id1(category_id) - id2(tag_id)
// id1(article_id) - id2(tag_id)
// id1(article_id) - id2(archive_id)
type Corelation struct {
	global.BB_MODEL
	ID1  uint   `json:"id_1"` // article
	ID2  uint   `json:"id_2"` // tag || category
	Str1 string `gorm:"size:255" json:"str1"`
	Str2 string `gorm:"size:255" json:"str2"`
	Str3 string `gorm:"size:255" json:"str3"`
	Str4 string `gorm:"size:255" json:"str4"`
	Int1 int    `json:"int1"`
	Int2 int    `json:"int2"`
	Int3 int    `json:"int3"`
	Int4 int    `json:"int4"`
	Type int    `json:"type"` // type
}
