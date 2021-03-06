package search

import (
	"fmt"
	"log"
)

// Result 搜索结果
type Result struct {
	Field   string
	Content string
}

// Matcher 定义要实现的新搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 为每个数据源单独启动goroutine来执行这个函数
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

// Display 显示查询结果
func Display(results chan *Result) {
	// 通道会一直被阻塞，直到有结果写入
	// 一旦通道被关闭，for循环就会终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
