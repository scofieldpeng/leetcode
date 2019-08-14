package flipAndInvertImage

// TODO
func flipAndInvertImage(A [][]int) [][]int {
	for k, v := range A {
		subLen := len(v)
		for i := 0; i < (len(v)+1)/2; i++ {
			A[k][i], A[k][subLen-i-1] = A[k][subLen-i-1]^1, A[k][i]^1
		}
	}

	return A
}
