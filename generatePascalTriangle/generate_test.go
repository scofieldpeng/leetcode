package generatePascalTriangle

import (
	"fmt"
	"testing"
)

func Test_generate(t *testing.T) {
	fmt.Println(generate(1))
	fmt.Println(generate(2))
	fmt.Println(generate(3))
}
