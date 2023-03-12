package models

import (
	"github.com/jinzhu/gorm"
	"mywork/pkg/model"
)

type LiteBase struct {
	gorm.Model
	Content string `gorm:"type:text;null"`     //存放数据
	Type    string `gorm:"type:int;default:0"` //存放数据
}

/**
创建数据
*/
func UpdateBase(content string) {
	DeleteBase("1")
	var base LiteBase
	base.Content = content
	base.Type = "1"
	model.Db.Save(&base)
}

/**
获取数据
*/
func GetAbort() (list LiteBase, err error) {

	return list, model.Db.Where("type = ?", 1).Order("created_at desc").Limit(1).Find(&list).Error
}

/**
获取数据
*/
func GetPlacard() (list LiteBase, err error) {

	return list, model.Db.Where("type = ?", 2).Order("created_at desc").Limit(1).Find(&list).Error
}

/**
设置公告
*/
func UpdatePlacard(content string) {
	DeleteBase("2")
	var base LiteBase
	base.Content = content
	base.Type = "2"
	model.Db.Save(&base)
}

/**
删除
*/
func DeleteBase(str string) {

	var base LiteBase

	model.Db.Delete(&base).Where("type = ?", str)

}
