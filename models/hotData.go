package models

import (
	"buzzGen/helpers"
	"fmt"
)

type TblHotData struct {
	Id         int    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL;comment:'主键'"`
	OriId      int    `gorm:"column:ori_id;NOT NULL;comment:'ori_id'"`
	OrderBy    int    `gorm:"column:order_by;NOT NULL;comment:'顺序'"`
	Category   string `gorm:"column:category;NOT NULL;comment:'类别'"`
	CreateTime int32  `gorm:"column:create_time;NOT NULL;comment:'创建时间'"`
	UpdateTime int32  `gorm:"column:update_time;NOT NULL;comment:'更新时间'"`
}

func (*TblHotData) TableName() string {
	return "tblHotData"
}

func (m *TblHotData) Insert() error {
	db := helpers.MySQLClient
	sqlDb, err := db.DB()
	err = sqlDb.Ping()
	if err != nil {
		err, db = helpers.MysqlReconnect()
		if err != nil {
			helpers.BuzzLogger.Error(fmt.Sprintf("mysql reconnect error: %v", err))
			return err
		}
	}
	if err := db.Create(m).Error; err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("insert data error: %v", err))
		return err
	}
	return nil
}

func DelHotDataByCategory(category string) error {
	db := helpers.MySQLClient
	err := db.Where("category = ?", category).Delete(&TblHotData{}).Error
	if err != nil {
		helpers.BuzzLogger.Error("del hot by category failed " + err.Error())
		return err
	}
	return nil
}
func AddHotDataList(list []TblHotData) error {
	db := helpers.MySQLClient
	err := db.Create(&list).Error
	if err != nil {
		helpers.BuzzLogger.Error("add hot data list failed " + err.Error())
		return err
	}
	return nil
}

func GetHotDataListByCategory(category string) (list []TblHotData, err error) {
	db := helpers.MySQLClient
	err = db.Where("category = ?", category).Find(&list).Error
	if err != nil {
		helpers.BuzzLogger.Error("get hot data list by category failed " + err.Error())
		return
	}
	return
}
