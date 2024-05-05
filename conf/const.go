package conf

type DataSourceRet struct {
	Url          string `json:"Url"`
	HeadImageUrl string `json:"HeadImageUrl"`
	OriTitle     string `json:"OriTitle"`
	Category     string `json:"Category"`
	ParsedData   string `json:"ParsedData"`
	OriLang      string `json:"OriLang"`
	FetchDetail  string `json:"FetchDetail"`
	OriOrderBy   int    `json:"OriOrderBy"`
	NewsTime     int    `json:"NewsTime"`
	FetchTime    int    `json:"FetchTime"`
}
