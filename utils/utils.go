package utils

func Findconsquetive(m map[int]int) (keys []int) {
	for k, _ := range m {
		if _, ok := m[k+1]; ok {
			keys = append(keys, k)
		}
	}
	return
}
func Sort(m []int) (keys []int) {
	for i := range m {
		for j := i + 1; j < len(m); j++ {
			if m[i] > m[j] {
				m[i], m[j] = m[j], m[i]
			}
		}
	}
	return m
}
