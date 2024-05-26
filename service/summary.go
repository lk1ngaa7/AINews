package service

import (
	"buzzGen/constant"
	"buzzGen/helpers"
	"buzzGen/models"
	"fmt"
)

func HandleSummary(cate string) {
	hotList, err := models.GetHotDataListByCategory("hn")
	if err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("GetHotDataListByCategory Error: %v", err))
		return
	}
	for _, hot := range hotList {
		oriData := models.GetOriDataById(hot.OriId)
		if oriData.ID == 0 {
			continue
		}
		summData, err := models.GetZhDataByOriId(oriData.ID)
		if summData.ID > 0 {
			continue
		}
		err, sum, title, translate := dealByGpt(oriData)
		if err != nil {
			helpers.BuzzLogger.Error(fmt.Sprintf("Insert Error: %v", err))
			continue
		}
		summData.Summary = sum
		summData.Title = title
		summData.TransText = translate
		summData.OriID = oriData.ID
		summData.Status = "succ"
		err = summData.Insert()
		if err != nil {
			helpers.BuzzLogger.Error(fmt.Sprintf("Insert Error: %v", err))
			continue
		}
	}
}

func dealByGpt(oriData models.TblOriData) (err error, sum string, title string, translateStr string) {
	var qStr string
	//1.生成总结
	qStr = fmt.Sprintf("请根据文章内容给我生产文章总结。\n 要求：1.总结的内容不能超过200个汉字字符。2.总结的内容需要尽可能的符合原文表达的含义同时符合中文的阅读习惯。3.总结结果只包含总结的内容，不能有任何其他字符！ \n 文章内容是:%s。", oriData.ParsedData)
	ret, err := helpers.LLMDeepSeek(qStr, constant.SYSTEM_ZH_SUMMARY_PROMPT)
	if err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("LLMDeepSeek Error: %v", err))
		return
	}
	//2.生成标题
	qStr = fmt.Sprintf("请将我给的文章标题翻译成中文 \n 要求：1.翻译可以参考我给出的中文文章摘要总结但是必须翻译出原来文章标题的实际含义。2.翻译结果只包含翻译后的标题内容，不能有任何其他字符！ \n 文章标题是:%s，文章摘要总结是:%s。", oriData.OriTitle, ret)
	retTitle, err := helpers.LLMDeepSeek(qStr, constant.SYSTEM_ZH_TRANS_PROMPT)
	if err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("LLMDeepSeek Error: %v", err))
		return
	}
	//3.生成翻译
	qStr = fmt.Sprintf("请将我给的文章翻译成中文 \n 要求：1.翻译需要保持原文的含义。2.翻译结果只包含翻译后的内容，不能有任何其他字符！ \n 文章原文内容是:%s。", oriData.ParsedData)
	translate, err := helpers.LLMDeepSeek(qStr, constant.SYSTEM_ZH_TRANS_PROMPT)
	if err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("LLMDeepSeek Error: %v", err))
		return
	}
	helpers.BuzzLogger.Info(fmt.Sprintf("Summary: %s, Title: %s, Translate: %s", ret, retTitle, translate))
	return nil, ret, retTitle, translate
}
