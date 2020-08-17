package strsim

import (
	"math"

	"github.com/go-ego/gse"
)

var seg = gse.New("zh", "alpha")

func vec(txtA, txtB string) (vecA, vecB []float64) {
	la := seg.Cut(txtA, true)
	la = seg.Trim(la)
	lb := seg.Cut(txtB, true)
	lb = seg.Trim(lb)

	ma := map[string]int{}
	mb := map[string]int{}
	mAll := map[string]int{}

	getCnt := func(l []string, m map[string]int) {
		for _, s := range l {
			if cnt, has := m[s]; has {
				m[s] = cnt + 1
			} else {
				m[s] = 1
			}
			if cnt, has := mAll[s]; has {
				mAll[s] = cnt + 1
			} else {
				mAll[s] = 1
			}
		}
	}

	getCnt(la, ma)
	getCnt(lb, mb)

	for s, cnt := range mAll {
		f, _ := seg.Find(s)
		mAll[s] = cnt + f
	}

	var lAll []string
	for s := range mAll {
		lAll = append(lAll, s)
	}
	getT := func(m map[string]int) []float64 {
		ret := make([]float64, len(lAll))
		for i, s := range lAll {
			ret[i] = float64(m[s]) / float64(mAll[s])
		}
		return ret
	}

	vecA = getT(ma)
	vecB = getT(mb)
	return vecA, vecB
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

func TfidfLoadDict(files ...string) error {
	return seg.LoadDict()
}

func Tfidf(txtA, txtB string) float64 {
	vecA, vecB := vec(txtA, txtB)
	return cosine(vecA, vecB)
}
