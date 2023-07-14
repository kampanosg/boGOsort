package bogosort

func Sort(s []any) []any {
	if len(s) < 2 {
		return s
	}
	return nil
}

func Shuffle(s []int) []int {
	return nil
}

func IsSorted(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false
		}
	}
	return true
}
