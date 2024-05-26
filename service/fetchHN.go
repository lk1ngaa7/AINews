package service

import (
	"buzzGen/data"
	"buzzGen/helpers"
	"buzzGen/models"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type HNFetchDetail struct {
	ImgList []string `json:"imgList"`
	HnId    int      `json:"hnId"`
}

func FetchHnData() (err error) {
	listId, err := data.GetTopHNList()
	if err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("Error: %v", err))
		return err
	}
	idMapList, err := data.GetBuzzHNDetail()
	var newHotList []models.TblHotData
	for index, id := range listId {
		key := "https://news.ycombinator.com/item?id=" + strconv.Itoa(id)
		hnd, err := data.GetHnDetail(id)
		if err != nil {
			helpers.BuzzLogger.Error(fmt.Sprintf("Error: %v", err))
			continue
		}
		if hnd.Url == "" {
			hnd.Url = key
		}
		oriData, err := models.GetOriDataByUrl(hnd.Url)
		if err != nil {
			helpers.BuzzLogger.Error(fmt.Sprintf("Error: %v", err))
		}
		if oriData.ID != 0 {
			helpers.BuzzLogger.Info(fmt.Sprintf("url: %v is already exist", hnd.Url))
			oriData.OriOrderBy = index
			err = oriData.Update()
			if err != nil {
				helpers.BuzzLogger.Error(fmt.Sprintf("update ori data Error: %v", err))
			}
			tmpHot := models.TblHotData{
				OriId:      oriData.ID,
				OrderBy:    index,
				Category:   "hn",
				CreateTime: int32(time.Now().Unix()),
				UpdateTime: int32(time.Now().Unix()),
			}
			newHotList = append(newHotList, tmpHot)
			if len(newHotList) == 10 {
				break
			}
			continue
		}
		markdown, err := data.GetMarkDownByJina(hnd.Url)
		if err != nil {
			helpers.BuzzLogger.Error(fmt.Sprintf("jina Error: %v", err))
			continue
		}
		image := ""
		if _, ok := idMapList[key]; ok {
			image = idMapList[key]
		}
		var mdImdList []string
		if image == "" {
			mdImdList = helpers.GetMdImage(markdown)
			if len(mdImdList) > 0 {
				image = mdImdList[0]
			}
			//取mdImgList 前三个
			if len(mdImdList) > 3 {
				mdImdList = mdImdList[:3]
			}
		}
		fetD, err := json.Marshal(HNFetchDetail{
			ImgList: mdImdList,
			HnId:    id,
		})
		tblOriData := models.TblOriData{
			OriOrderBy:   index, //使用接口的顺序和 hackernews 保持一直
			Url:          hnd.Url,
			FetchTime:    int(time.Now().Unix()),
			HeadImageUrl: image,
			OriTitle:     hnd.Title,
			ParsedData:   markdown,
			Category:     "hackerNews",
			OriLang:      "en",
			NewsTime:     hnd.Time,
			FetchDetail:  string(fetD),
			IsDeleted:    0,
			ExtStr:       strconv.Itoa(id),
		}
		err = tblOriData.Insert()
		if err != nil {
			helpers.BuzzLogger.Error(fmt.Sprintf("Insert Error: %v", err))
		} else {
			tmpHot := models.TblHotData{
				OriId:      tblOriData.ID,
				OrderBy:    index,
				Category:   "hn",
				CreateTime: int32(time.Now().Unix()),
				UpdateTime: int32(time.Now().Unix()),
			}
			newHotList = append(newHotList, tmpHot)
			if len(newHotList) == 10 {
				break
			}
		}
	}
	//如果newHotList长度大于10.只保留10个
	if len(newHotList) > 10 {
		newHotList = newHotList[:10]
	}
	err = DelAndAddHotData("hn", newHotList)
	if err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("DelAndAddHotData Error: %v", err))
	}
	return
}
