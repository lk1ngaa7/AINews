package models

import (
	"buzzGen/helpers"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type TblOriData struct {
	ID           int    `gorm:"column:id;primary_key;AUTO_INCREMENT"` // 主键
	Url          string `gorm:"column:url"`                           // url
	HeadImageUrl string `gorm:"column:head_image_url"`                // 头图
	OriTitle     string `gorm:"column:ori_title"`                     // 原始标题
	Category     string `gorm:"column:category"`                      // 类别
	ParsedData   string `gorm:"column:parsed_data"`                   // 抓取后的文本数据
	OriLang      string `gorm:"column:ori_lang"`                      // 原始语种
	FetchDetail  string `gorm:"column:fetch_detail"`                  // 原始数据
	ExtStr       string `gorm:"column:ext_str"`                       // 扩展数据，不同数据源不一定
	OriOrderBy   int    `gorm:"column:ori_order_by;NOT NULL"`         // 原始排序 有的是时间有的是点赞数
	NewsTime     int    `gorm:"column:news_time;NOT NULL"`            // 新闻时间
	FetchTime    int    `gorm:"column:fetch_time;NOT NULL"`           // 抓取时间
	IsDeleted    int    `gorm:"column:is_deleted;default:0;NOT NULL"` // 0-未删除 1-已删除
}

func (*TblOriData) TableName() string {
	return "tblOriData"
}

// update
func (m *TblOriData) Update() error {
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
	if err := db.Save(m).Error; err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("update data error: %v", err))
		return err
	}
	return nil
}
func (m *TblOriData) Insert() error {
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
func GetOriDataByUrl(url string) (data TblOriData, err error) {
	db := helpers.MySQLClient
	if err = db.Where("url = ?", url).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.BuzzLogger.Info(fmt.Sprintf("get data empty: %v", err))
			return data, nil
		}
		helpers.BuzzLogger.Error(fmt.Sprintf("get data error: %v", err))
		return
	}
	return
}
