package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(&a).Updates(values).Where("id =? AND is_delete = ?", a.ID, 0).Error
	if err != nil {
		return err
	}
	return nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND state = ? AND is_delete = ?", a.ID, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	err := db.Where("id = ? AND is_delete = ?", a.Model.ID, 0).Delete(&a).Error
	if err != nil {
		return err
	}
	return nil
}

type ArticleRow struct {
	ArticleID     uint32
	TagId         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id as article_id", "ar.title as article_title", "ar.desc as article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id as tag_id", "t.name as tag_name"}...)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id=t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar On at.article_id=ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_delete = ?", tagID, a.State, 0).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		row := &ArticleRow{}
		err := rows.Scan(&row.ArticleID, &row.ArticleTitle, &row.ArticleDesc, &row.CoverImageUrl, &row.Content, &row.TagId, &row.TagName)
		if err != nil {
			return nil, err
		}
		articles = append(articles, row)
	}
	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id=t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id=ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_delete = ?", tagID, a.State, 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
