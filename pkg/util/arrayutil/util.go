package arrayutil

func Contains[T string | int | int32 | int64](array []T, val T) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return true
		}
	}
	return false
}
