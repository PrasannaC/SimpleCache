package Utils

import "math"

/*
	The hash function to get int from key.
	Source: Wiki page for java.lang.String hash function.
 */
func GetStringHash(strKey string) uint64 {
	var hash uint64 = 0
	n := len(strKey) - 1
	for i, v := range strKey {
		hash += uint64(v) * uint64(math.Pow(31, float64(n-i)))
	}
	return hash
}
