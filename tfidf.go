package strsim

import (
	"math"

	"github.com/go-ego/gse"
)

var Segmenter = gse.New("zh", "alpha")

func str2vec(str string) map[string]int {
	l := Segmenter.Cut(str, true)
	l = Segmenter.Trim(l)
	m := map[string]int{}
	for _, s := range l {
		if cnt, has := m[s]; has {
			m[s] = cnt + 1
		} else {
			m[s] = 1
		}
	}
	return m
}

func cosine(vecA, vecB []float64) float64 {
	var product, squareSumA, squareSumB float64
	for i, v := range vecA {
		product += v * vecB[i]
		squareSumA += v * v
		squareSumB += vecB[i] * vecB[i]
	}

	if squareSumA == 0 || squareSumB == 0 {
		return 0
	}

	return product / (math.Sqrt(squareSumA) * math.Sqrt(squareSumB))
}

func NewTfidfCompare(baseStr string) func(str string) float64 {
	mb := str2vec(baseStr)
	sim := func(str string) float64 {
		m := str2vec(str)
		ma := map[string]int{}
		for s, cnt := range mb {
			ma[s] = cnt
		}
		for s, cnt := range m {
			ma[s] = ma[s] + cnt
		}
		var la []string
		for s := range ma {
			la = append(la, s)
		}
		vec := func(m map[string]int) []float64 {
			ret := make([]float64, len(la))
			for i, s := range la {
				ret[i] = float64(m[s]) / float64(ma[s])
			}
			return ret
		}
		va := vec(mb)
		vb := vec(m)
		return cosine(va, vb)
	}
	return sim
}
