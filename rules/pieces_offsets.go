package rules

var (
	kingMoveOffsets = []square{
		{file: -1, rank: 1}, {file: 0, rank: 1}, {file: 1, rank: 1},
		{file: -1, rank: 0} /*****************/, {file: 1, rank: 0},
		{file: -1, rank: -1}, {file: 0, rank: -1}, {file: 1, rank: -1},
	}

	knightMoveOffsets = []square{
		{file: -1, rank: 2}, {file: -1, rank: -2},
		{file: 1, rank: 2}, {file: 1, rank: -2},
		{file: -2, rank: 1}, {file: -2, rank: -1},
		{file: 2, rank: 1}, {file: 2, rank: -1},
	}

	rookMoveOffsetLines = [][]square{
		{
			{rank: 1}, {rank: 2}, {rank: 3}, {rank: 4}, {rank: 5}, {rank: 6}, {rank: 7},
		},
		{
			{rank: -1}, {rank: -2}, {rank: -3}, {rank: -4}, {rank: -5}, {rank: -6}, {rank: -7},
		},
		{
			{file: 1}, {file: 2}, {file: 3}, {file: 4}, {file: 5}, {file: 6}, {file: 7},
		},
		{
			{file: -1}, {file: -2}, {file: -3}, {file: -4}, {file: -5}, {file: -6}, {file: -7},
		},
	}

	bishopMoveOffsetLines = [][]square{
		{
			{rank: 1, file: 1}, {rank: 2, file: 2}, {rank: 3, file: 3}, {rank: 4, file: 4},
			{rank: 5, file: 5}, {rank: 6, file: 6}, {rank: 7, file: 7},
		},
		{
			{rank: -1, file: 1}, {rank: -2, file: 2}, {rank: -3, file: 3}, {rank: -4, file: 4},
			{rank: -5, file: 5}, {rank: -6, file: 6}, {rank: -7, file: 7},
		},
		{
			{rank: 1, file: -1}, {rank: 2, file: -2}, {rank: 3, file: -3}, {rank: 4, file: -4},
			{rank: 5, file: -5}, {rank: 6, file: -6}, {rank: 7, file: -7},
		},
		{
			{rank: -1, file: -1}, {rank: -2, file: -2}, {rank: -3, file: -3}, {rank: -4, file: -4},
			{rank: -5, file: -5}, {rank: -6, file: -6}, {rank: -7, file: -7},
		},
	}

	queenMoveOffsetLines = append(rookMoveOffsetLines, bishopMoveOffsetLines...)

	whitePawnAdvancementOffsets = []square{{file: 0, rank: 1}}
	blackPawnAdvancementOffsets = []square{{file: 0, rank: -1}}

	whitePawnInitialAdvancementOffsets = append(whitePawnAdvancementOffsets, square{file: 0, rank: 2})
	blackPawnInitialAdvancementOffsets = append(blackPawnAdvancementOffsets, square{file: 0, rank: -2})

	whitePawnCaptureOffsets = []square{
		{file: -1, rank: 1},
		{file: 1, rank: 1},
	}
	blackPawnCaptureOffsets = []square{
		{file: -1, rank: -1},
		{file: 1, rank: -1},
	}
)
