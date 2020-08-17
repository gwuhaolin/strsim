package strsim

import (
	"testing"
)

func TestTfidf(t *testing.T) {
	textA := `KastKing路亚竿套装全套远投枪柄水滴轮超硬碳素杆抛竿海竿钓鱼竿`
	textB := `绑好台钓成品方便线组套装全套鲫鱼钓鱼线钩子线夹八字环渔具用品`
	textC := `迪卡侬海竿路亚竿矶钓竿超轻超硬钓鱼竿手海两用竿长节CAP`
	textD := `KastKing富士碳素直柄路亚竿套装全套纺车轮远投UL马口竿ML钓鱼竿`
	textE := `路亚竿套装全套远投枪柄水滴轮超硬碳素杆抛竿海竿钓鱼竿KastKing牌`
	s1 := Tfidf(textA, textB)
	t.Log(s1)
	s2 := Tfidf(textA, textC)
	t.Log(s2)
	s3 := Tfidf(textA, textD)
	t.Log(s3)
	s4 := Tfidf(textA, textE)
	t.Log(s4)
}

func TestCosine(t *testing.T) {
	c1 := cosine([]float64{1, 0.01}, []float64{2, 0.01})
	c2 := cosine([]float64{1, 0.01}, []float64{3, 0.01})

	c3 := cosine([]float64{1, 0.05}, []float64{2, 0.05})
	c4 := cosine([]float64{1, 0.05}, []float64{3, 0.05})

	c5 := cosine([]float64{1, 1}, []float64{2, 1})
	c6 := cosine([]float64{1, 1}, []float64{3, 1})
	t.Log(c1, c2, c3, c4, c5, c6)
}
