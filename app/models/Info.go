package models

import (
	. "tattoo_code/database"

	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null"`
	Content string `gorm:"type:varchar(100);not null"`
}

// CreateInfo 创建信息
func (info *Info) CreateInfo() error {
	reslut := DB.Create(&info)
	if reslut.Error != nil {
		return reslut.Error
	}
	return nil
}

// 根据传入的id获取信息，返回数据信息或者错误
func (info *Info) GetInfoById(id int) (Info, error) {
	var info1 Info
	reslut := DB.First(&info1, id)
	if reslut.Error != nil {
		return info1, reslut.Error
	}
	return info1, nil
}

// 修改信息
func (info *Info) UpdateInfo() error {
	reslut := DB.Save(&info)
	if reslut.Error != nil {
		return reslut.Error
	}
	return nil
}

// 删除信息
func (info *Info) DeleteInfo() error {
	reslut := DB.Delete(&info)
	if reslut.Error != nil {
		return reslut.Error
	}
	return nil
}
