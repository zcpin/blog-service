package dao

import (
	"github.com/zcpin/blog-service/internal/model"
	"github.com/zcpin/blog-service/pkg/app"
)

func (d *Dao) GetTag(tagID uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{State: state, Model: &model.Model{ID: tagID}}
	return tag.Get(d.engine)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{Name: name, State: state, Model: &model.Model{CreatedBy: createdBy}}
	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			ID:         id,
			ModifiedBy: modifiedBy,
		},
	}
	value := map[string]interface{}{
		"modified_by": modifiedBy,
		"state":       state,
	}
	if name != "" {
		value["name"] = name
	}
	return tag.Update(d.engine, value)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}
