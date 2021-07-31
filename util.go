package strsim

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func avg(source []float64) (avg float64) {
	for _, f := range source {
		avg += f
	}
	return avg / float64(len(source))
}
