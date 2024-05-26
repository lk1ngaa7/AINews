package service

import (
	"buzzGen/models"
)

func DelAndAddHotData(category string, hotList []models.TblHotData) error {
	err := models.DelHotDataByCategory(category)
	if err != nil {
		return err
	}
	err = models.AddHotDataList(hotList)
	if err != nil {
		return err
	}
	return nil
}
