package server

import (
	"context"
	"encoding/json"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/olivere/elastic/v7"
	"log"
	"regexp"
	"strings"
)

type SearchServer struct {
}

func (s *SearchServer) Search(query string, from, size, filterType, filterGroup int) (any, error) {
	// 创建针对 title 和 content 的匹配查询
	matchTitleQuery := elastic.NewMatchPhraseQuery("title", query)
	matchContentQuery := elastic.NewMatchPhraseQuery("content", query)

	// 构建布尔查询的主体
	boolQuery := elastic.NewBoolQuery()

	// 添加 title 和 content 的匹配条件
	boolQuery.Must(elastic.NewBoolQuery().Should(matchTitleQuery, matchContentQuery))

	// 如果提供了类型过滤器，则添加到查询中
	if filterType != -1 { // 假设 -1 表示不过滤此字段
		boolQuery.Must(elastic.NewTermQuery("type", filterType))
	}

	// 如果提供了组过滤器，则添加到查询中
	if filterGroup != -1 { // 假设 -1 表示不过滤此字段
		boolQuery.Must(elastic.NewTermQuery("groupId", filterGroup))
	}

	// 设置高亮，仅对 content 字段
	highlight := elastic.NewHighlight().
		Field("content").         // 高亮content字段
		Field("title").           // 高亮title字段
		PreTags("<em>").          // 设置高亮开始标签
		PostTags("</em>").        // 设置高亮结束标签
		FragmentSize(150).        // 设置每个高亮片段的字符长度
		NumOfFragments(3).        // 设置每个字段返回的最大高亮片段数
		HighlightQuery(boolQuery) // 设置高亮查询与主查询相同

	// 执行搜索，并加入分页参数和高亮设置
	searchResult, err := db.EDB.Search().
		Index(config.SettingsConfig.Es.Index).
		Query(boolQuery).
		Highlight(highlight).
		From(from).Size(size).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		log.Fatalf("Error searching: %s", err)
		return nil, err
	}

	// 创建结果数组
	var results []models.Search

	for _, hit := range searchResult.Hits.Hits {
		var s models.Search
		err := json.Unmarshal(hit.Source, &s)
		if err != nil {
			// 处理错误
		}

		// 提取高亮结果
		if hit.Highlight != nil {
			if highlightContent, ok := hit.Highlight["content"]; ok {
				// 合并相邻的高亮标签
				s.Content = mergeAdjacentTags(strings.Join(highlightContent, " "))
			}
			if highlightTitle, ok := hit.Highlight["title"]; ok {
				s.Title = mergeAdjacentTags(strings.Join(highlightTitle, " "))
			}
		}

		results = append(results, s)
	}

	return map[string]any{
		"list":  results,
		"total": searchResult.TotalHits(),
	}, nil
}

// mergeAdjacentTags 合并相邻的高亮标签
func mergeAdjacentTags(text string) string {
	// 使用正则表达式来查找并合并相邻的高亮标签
	re := regexp.MustCompile(`</em><em>`)
	return re.ReplaceAllString(text, "")
}
