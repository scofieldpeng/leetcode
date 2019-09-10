package numSmallerByFrequency

import (
	"fmt"
	"testing"
)

func Test_numSmallerByFrequency(t *testing.T) {
	var a = []string{"bba", "abaaaaaa", "aaaaaa", "bbabbabaab", "aba", "aa", "baab", "bbbbbb", "aab", "bbabbaabb"}
	var b = []string{"aaabbb", "aab", "babbab", "babbbb", "b", "bbbbbbbbab", "a", "bbbbbbbbbb", "baaabbaab", "aa"}

	var res = numSmallerByFrequency(a,b)
	fmt.Println(res)
}
