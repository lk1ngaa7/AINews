package data

import (
	"buzzGen/helpers"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func GetMarkDownByJina(url string) (str string, err error) {
	client := resty.New()
	res, err := client.R().Get("https://r.jina.ai/" + url)
	str = string(res.Body())
	if err != nil {
		for i := 0; i < 2; i++ {
			res, err = client.R().Get("https://r.jina.ai/" + url)
			str = string(res.Body())
			if helpers.IsJSON(str) {
				return "", fmt.Errorf("jina failed")
			}
			if err == nil && len(str) > 0 {
				helpers.BuzzLogger.Info(fmt.Sprintf("GetMarkDownByJina retry success: %v", i))
				return
			}
		}
		helpers.BuzzLogger.Error(fmt.Sprintf(" GetMarkDownByJina Error: %v", err))
		return
	}
	return
}
