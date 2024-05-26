package models

import (
	"buzzGen/helpers"
	"fmt"
)

type TblZhData struct {
	ID         int    `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 主键
	OriID      int    `gorm:"column:ori_id;NOT NULL"`               // ori_id
	Status     string `gorm:"column:status"`                        // 解析状态 succ ，pending，failed
	Title      string `gorm:"column:title"`                         // 翻译后的标题
	Summary    string `gorm:"column:summary"`                       // 总结数据
	TransText  string `gorm:"column:trans_text"`                    // 翻译结果
	CreateTime int    `gorm:"column:create_time;NOT NULL"`          // 创建时间
	UpdateTime int    `gorm:"column:update_time;NOT NULL"`          // 更新时间
	IsDeleted  int    `gorm:"column:is_deleted;default:1;NOT NULL"` // 0-未删除 1-已删除
}

func (*TblZhData) TableName() string {
	return "tblZhData"
}

func (m *TblZhData) Insert() error {
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

func GetZhDataByOriId(oriId int) (data TblZhData, err error) {
	db := helpers.MySQLClient
	err = db.Where("ori_id = ?", oriId).First(&data).Error
	if err != nil {
		helpers.BuzzLogger.Error("get zh data by ori_id failed " + err.Error())
		return
	}
	return

}
