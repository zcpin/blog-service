package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zcpin/blog-service/global"
	"github.com/zcpin/blog-service/pkg/setting"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreateBy   string `json:"create_by"`
	CreateOn   string `json:"create_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn string `json:"modified_on"`
	DeleteOn   uint32 `json:"delete_on"`
	IsDelete   uint8  `json:"is_delete"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(databaseSetting.DBType, dsn)
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunModel == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdelConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}
