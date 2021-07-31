# 字符串相似性计算

## TFIDF

基于分词+统计概率

```go
package main

import "github.com/gwuhaolin/strsim"

func main() {
	sim := strsim.NewTfidfCompare(`KastKing路亚竿套装全套远投枪柄水滴轮超硬碳素杆抛竿海竿钓鱼竿`)
	textB := `绑好台钓成品方便线组套装全套鲫鱼钓鱼线钩子线夹八字环渔具用品`
	textC := `迪卡侬海竿路亚竿矶钓竿超轻超硬钓鱼竿手海两用竿长节CAP`
	s1, s2 := sim(textB), sim(textC)
	println(s1, s2)
}
```

分词基于[gse](http://github.com/go-ego/gse)，自定义分词词典：

```go

```

## Jaro

基于编辑距离

