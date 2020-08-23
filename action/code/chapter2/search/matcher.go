package search

import (
	"log"
)

//Result保存搜索结果
type Result struct {
	Field   string
	Content string
}

//Matcher定义了要实现的
//新搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}
