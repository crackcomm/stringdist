package stringdist

import (
	"strings"

	"github.com/crackcomm/stringdist/levenshtein"
)

// Join - Joins two strings slices using levenshtein comparision with given
// distance as a treshold to treat two values as same values.
func Join(dst, src []string, maxDist int) (result []string) {
	result = dst
	for _, value := range src {
		if len(value) == 0 {
			continue
		}
		if !containsString(result, value, maxDist) {
			result = append(result, value)
		}
	}
	return
}

func containsString(list []string, str string, maxDist int) (res bool) {
	if len(list) == 0 {
		return false
	}
	str = strings.ToLower(str)
	for _, value := range list {
		v := strings.ToLower(value)
		if maxDist == 0 && v == str {
			return true
		} else if _, has := levenshtein.DistanceThreshold(v, str, maxDist); has {
			return true
		}
	}
	return
}
