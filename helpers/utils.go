package helpers

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// IsJSON 判断字符串是否是有效的JSON
func IsJSON(s string) bool {
	var js map[string]interface{}
	// 尝试将字符串解析为JSON对象
	return json.Unmarshal([]byte(s), &js) == nil
}

// ExtractJSON 从给定字符串中提取JSON字符串
func ExtractJSON(str string) (string, error) {
	// 定义正则表达式模式,支持多行JSON
	pattern := `(?s)\{.*?\}`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(str, -1)

	// 如果找到多个匹配项,返回最长的那个
	if len(matches) > 0 {
		longestMatch := ""
		for _, match := range matches {
			if len(match) > len(longestMatch) {
				longestMatch = match
			}
		}
		return longestMatch, nil
	}

	// 如果正则匹配失败,尝试使用字符串切片和计数括号的方式
	braces := 0
	start := 0
	for i, c := range str {
		switch c {
		case '{':
			if braces == 0 {
				start = i
			}
			braces++
		case '}':
			braces--
			if braces == 0 {
				return str[start : i+1], nil
			}
		}
	}

	// 未找到JSON字符串
	return "", fmt.Errorf("no valid JSON string found in '%s'", str)
}
