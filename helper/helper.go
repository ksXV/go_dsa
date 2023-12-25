package helper

func CompareTwoUniDimensionalArrays[T comparable](arrOne []T, arrTwo []T) bool {
	if len(arrOne) != len(arrTwo) {
		return false
	}
	for i := 0; i < len(arrOne); i++ {
		if arrOne[i] != arrTwo[i] {
			return false
		}
	}
	return true
}
