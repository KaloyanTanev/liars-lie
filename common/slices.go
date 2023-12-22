package common

func MostFrequentInteger(slice []int) int {
	m := map[int]int{}
	var maxCnt int
	var freq int
	for _, a := range slice {
		m[a]++
		if m[a] > maxCnt {
			maxCnt = m[a]
			freq = a
		}
	}
	return freq
}

func SmallestInteger(slice []int) int {
	smallest := slice[0]
	for _, num := range slice[1:] {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}

func RemoveIndex[T any](slice []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}

func GetDistinctIntegers(slice []int) map[int]int {
	m := make(map[int]int)
	for _, s := range slice {
		if val, ok := m[s]; ok {
			m[s] = val + 1
		} else {
			m[s] = 1
		}
	}
	return m
}
