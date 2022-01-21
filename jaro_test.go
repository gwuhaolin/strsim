package strsim

import "testing"

func TestJaroMatterPrefix(t *testing.T) {
	t.Log(JaroMatterPrefix(`锄头湖`, `黄土下湖`))
	t.Log(JaroMatterPrefix(`金井水库`, `金泰垂钓`))

	t.Log(JaroMatterPrefix(
		`KastKing路亚竿套装全套远投枪柄水滴轮超硬碳素杆抛竿海竿钓鱼竿`,
		`KastKing路亚竿套装全套远投枪柄水滴轮超硬碳素杆抛竿海竿钓鱼竿`))
	t.Log(JaroMatterPrefix(
		`KastKing路亚竿套装全套远投枪柄水滴轮超硬碳素杆抛竿海竿钓鱼竿`,
		`KastKing富士碳素直柄路亚竿套装全套纺车轮远投UL马口竿ML钓鱼竿`))
	t.Log(JaroMatterPrefix(
		`KastKing路亚竿套装全套远投枪柄水滴轮超硬碳素杆抛竿海竿钓鱼竿`,
		`迪卡侬海竿路亚竿矶钓竿超轻超硬钓鱼竿手海两用竿长节CAP`))
}
