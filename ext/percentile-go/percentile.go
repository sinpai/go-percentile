package main

import (
	"C"
	"math"
	"sort"
)

func main() {}

//export percentile
func percentile(input []float64, percent float64) (percentile float64) {
	if len(input) == 0 || percent == 0 {
		return math.NaN()
	}
	c := sortedCopy(input)
	index := (percent / 100) * float64(len(c))
	if index == float64(int64(index)) {
		i := float64ToInt(index)
		percentile = mean([]float64{c[i-1], c[i]})
	} else if index >= 1 {
		i := float64ToInt(index)
		percentile = c[i-1]
	} else {
		return math.NaN()
	}
	return percentile
}
func float64ToInt(input float64) (output int) {
	r := round(input)
	return int(r)
}
func copyslice(input []float64) []float64 {
	s := make([]float64, len(input))
	copy(s, input)
	return s
}
func sortedCopy(input []float64) (copy []float64) {
	if sort.Float64sAreSorted(input) {
		return input
	}
	copy = copyslice(input)
	sort.Float64s(copy)
	return
}
func mean(input []float64) float64 {
	sum := sum(input)
	return sum / float64(len(input))
}
func sum(input []float64) (sum float64) {
	for _, n := range input {
		sum += n
	}
	return sum
}
func round(input float64) (rounded float64) {
	_, decimal := math.Modf(input)

	if decimal >= 0.5 {
		rounded = math.Ceil(input)
	} else {
		rounded = math.Floor(input)
	}
	return rounded
}
