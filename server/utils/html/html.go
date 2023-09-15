package html

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"unicode/utf8"
)

// GetHtmlSummary 获取html内容
func GetHtmlSummary(content string) (string, error) {
	reader, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return "", errors.New("parse_html_error")
	}
	innerText := ""
	reader.Find("p").Each(func(index int, item *goquery.Selection) {
		innerText += item.Text()
	})
	// 将字符串转换为 rune 切片
	runeText := []rune(innerText)
	// 确保不会越界，并且不会截断字符
	if utf8.RuneCountInString(innerText) > 128 {
		return string(runeText[:128]), nil
	} else {
		return innerText, nil
	}
}
