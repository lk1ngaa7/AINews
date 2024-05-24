package data

import (
	"buzzGen/helpers"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

type OneBuzzHN struct {
	Title             string    `json:"title"`
	Summary           string    `json:"summary"`
	ContentText       string    `json:"content_text"`
	Id                string    `json:"id"`
	Url               string    `json:"url"`
	DatePublished     time.Time `json:"date_published"`
	DateModified      time.Time `json:"date_modified"`
	OriginalPublished time.Time `json:"_original_published"`
	Score             int       `json:"_score"`
	NumComments       int       `json:"_num_comments"`
	Image             string    `json:"image"`
	Links             []struct {
		Url  string `json:"url"`
		Name string `json:"name"`
	} `json:"_links"`
}
type BuzzHNDetail struct {
	Title           string      `json:"title"`
	LatestBuildTime time.Time   `json:"_latest_build_time"`
	Language        string      `json:"language"`
	FeedUrl         string      `json:"feed_url"`
	Items           []OneBuzzHN `json:"items"`
}

func GetBuzzHNDetail() (idImgMap map[string]string, err error) {
	client := resty.New()
	var data BuzzHNDetail
	_, err = client.R().SetResult(&data).Get("https://hn.buzzing.cc/feed.json")
	if err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("Error: %v", err))
		return
	}
	idImgMap = make(map[string]string)
	helpers.BuzzLogger.Info(fmt.Sprintf("buzz news cnt: %v", len(data.Items)))
	for _, item := range data.Items {
		if len(item.Links) == 1 && item.Image != "" {
			idImgMap[item.Links[0].Url] = item.Image
		}
	}
	return
}

// 取前20条
func GetTopHNList() (data []int, err error) {
	client := resty.New()
	_, err = client.R().SetResult(&data).Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		helpers.BuzzLogger.Error(fmt.Sprintf("Error: %v", err))
		return
	}
	helpers.BuzzLogger.Info(fmt.Sprintf("top stories cnt: %v", len(data)))

	if len(data) > 100 {
		data = data[:20]
	}
	return
}

type HNDetail struct {
	Id    int    `json:"id"`
	Score int    `json:"score"`
	Time  int    `json:"time"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}

const (
	HNDETAIL_TYPE_STORY   = "story"
	HNDETAIL_TYPE_COMMENT = "comment"
	HNDETAIL_TYPE_JOB     = "job"
	HNDETAIL_TYPE_POLL    = "poll"
	HNDETAIL_TYPE_POLLOPT = "pollopt"
)

func GetHnDetail(itemId int) (hnd HNDetail, err error) {
	client := resty.New()
	_, err = client.R().SetResult(&hnd).Get("https://hacker-news.firebaseio.com/v0/item/" + strconv.Itoa(itemId) + ".json")
	if err != nil {
		helpers.BuzzLogger.Sugar().Warnw("getHnDetail err",
			"errMsg", err.Error(),
		)
		return
	}
	return
}
