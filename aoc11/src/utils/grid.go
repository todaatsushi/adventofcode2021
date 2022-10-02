package utils

func GetMoves() [8][2]int {
	var moves [8][2]int

	moves[0] = [2]int{1, 0}
	moves[1] = [2]int{-1, 0}

	moves[2] = [2]int{0, 1}
	moves[3] = [2]int{0, -1}

	moves[4] = [2]int{1, 1}
	moves[5] = [2]int{-1, -1}
	moves[6] = [2]int{1, -1}
	moves[7] = [2]int{-1, 1}

	return moves
}
