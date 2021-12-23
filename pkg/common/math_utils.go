package common

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Positive(x, y, z int) bool {
	return x >= 0 && y >= 0 && z >= 0
}
