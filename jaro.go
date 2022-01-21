package strsim

import (
	"github.com/xrash/smetrics"
	"math"
)

func unOrderArr(arr1, arr2 []string, compare func(str1, str2 string) float64) float64 {
	smallArr := arr1
	bigArr := arr2
	if len(arr1) > len(arr2) {
		smallArr = arr2
		bigArr = arr1
	}
	var maxSourceArr []float64
	var maxIndexArr []int
	for _, one := range smallArr {
		maxSource := compare(one, bigArr[0])
		maxIndex := 0
		for i := 1; i < len(bigArr); i++ {
			if contains(maxIndexArr, i) {
				continue
			}
			two := bigArr[i]
			x := compare(one, two)
			if x > maxSource {
				maxSource = x
				maxIndex = i
			}
		}
		maxSourceArr = append(maxSourceArr, maxSource)
		maxIndexArr = append(maxIndexArr, maxIndex)
	}
	return avg(maxSourceArr)
}

func Jaro(str1, str2 string) float64 {
	return smetrics.Jaro(str1, str2)
}

func JaroMatterPrefix(str1, str2 string) float64 {
	j := smetrics.Jaro(str1, str2)

	if j <= 0.7 {
		return j
	}

	arr1, arr2 := []rune{}, []rune{}
	for _, c := range str1 {
		arr1 = append(arr1, c)
	}
	for _, c := range str2 {
		arr2 = append(arr2, c)
	}
	prefixSize := int(math.Min(float64(len(arr1)), float64(len(arr2))))

	var prefixMatch float64
	for i := 0; i < prefixSize; i++ {
		if arr1[i] == arr2[i] {
			prefixMatch++
		} else {
			break
		}
	}

	return j + prefixMatch/float64(prefixSize)*(1.0-j)
}

// 找出两个无序数组的相似性
func JaroMatterPrefixUnOrderArr(arr1, arr2 []string) float64 {
	return unOrderArr(arr1, arr2, JaroMatterPrefix)
}

// 找出两个无序数组的相似性
func JaroUnOrderArr(arr1, arr2 []string) float64 {
	return unOrderArr(arr1, arr2, smetrics.Jaro)
}
