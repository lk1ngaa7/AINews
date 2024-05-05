package helpers

import (
	"regexp"

	"github.com/russross/blackfriday/v2"
)

func GetMdImage(md string) (imgList []string) {
	markdownBytes := []byte(md)
	htmlBytes := blackfriday.Run(markdownBytes)
	imgLinks := extractImageLinks(htmlBytes)

	for _, link := range imgLinks {
		imgList = append(imgList, link)
	}
	return
}

// 从 HTML 中提取图片链接
func extractImageLinks(html []byte) []string {
	var imgLinks []string

	// 在 HTML 中查找 img 标签
	imgRegex := regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)
	matches := imgRegex.FindAllSubmatch(html, -1)

	// 提取图片链接
	for _, match := range matches {
		if len(match) > 1 {
			imgLinks = append(imgLinks, string(match[1]))
		}
	}

	return imgLinks
}
