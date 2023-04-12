package utils

func IsCointainUint64(array []uint64, target uint64) bool {
	if array == nil {
		return false
	}
	for _, num := range array {
		if num == target {
			return true
		}
	}
	return false
}
