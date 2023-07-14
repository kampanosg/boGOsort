package bogosort

import (
	"math/rand"
	"time"
)

func Sort(s []int) {
	if len(s) < 2 {
		return
	}

	if IsSorted(s) {
		return
	}

	for {
		Shuffle(s)
		if IsSorted(s) {
			break
		}
	}
}

func Shuffle(s []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func IsSorted(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false
		}
	}
	return true
}
