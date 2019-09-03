package countCharacters

import "testing"

func Test_countCharacters(t *testing.T) {
	var a = []string{"cat", "bt", "hat", "tree"}
	var b = "attach"

	res := countCharacters(a, b)
	if res != 6 {
		t.Error("want 6,get 8")
	}
}
