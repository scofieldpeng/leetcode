package numRookCaptures

func numRookCaptures(board [][]byte) int {
	// first, find the location of rook
	rookX := 0
	rookY := 0

	num := 0

	for k, v := range board {
		for kk, vv := range v {
			if vv == 'R' {
				rookX = k
				rookY = kk
				break
			}
		}
		if rookX > 0 && rookY > 0 {
			break
		}
	}

	// west and east direction find
	for i := rookY; i > 0; i-- {
		if board[rookX][i] == 'p' {
			num++
			break
		}
		if board[rookX][i] == 'B' {
			break
		}
	}
	for i := rookY; i < 8; i++ {
		if board[rookX][i] == 'p' {
			num++
			break
		}
		if board[rookX][i] == 'B' {
			break
		}
	}

	// north and south direction find
	for i := rookX; i > 0; i-- {
		if board[i][rookY] == 'p' {
			num++
			break
		}
		if board[i][rookY] == 'B' {
			break
		}
	}
	for i := rookX; i < 8; i++ {
		if board[i][rookY] == 'p' {
			num++
			break
		}
		if board[i][rookY] == 'B' {
			break
		}
	}

	return num
}
