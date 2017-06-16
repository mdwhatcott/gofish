package rules

var (
	kingMoveOffsets = []Square{
		{File: -1, Rank: 1}, {File: 0, Rank: 1}, {File: 1, Rank: 1},
		{File: -1, Rank: 0} /*****************/, {File: 1, Rank: 0},
		{File: -1, Rank: -1}, {File: 0, Rank: -1}, {File: 1, Rank: -1},
	}

	rookMoveOffsetLines = [][]Square{
		{
			{Rank: 1}, {Rank: 2}, {Rank: 3}, {Rank: 4}, {Rank: 5}, {Rank: 6}, {Rank: 7},
		},
		{
			{Rank: -1}, {Rank: -2}, {Rank: -3}, {Rank: -4}, {Rank: -5}, {Rank: -6}, {Rank: -7},
		},
		{
			{File: 1}, {File: 2}, {File: 3}, {File: 4}, {File: 5}, {File: 6}, {File: 7},
		},
		{
			{File: -1}, {File: -2}, {File: -3}, {File: -4}, {File: -5}, {File: -6}, {File: -7},
		},
	}

	bishopMoveOffsetLines = [][]Square{
		{
			{Rank: 1, File: 1}, {Rank: 2, File: 2}, {Rank: 3, File: 3}, {Rank: 4, File: 4},
			{Rank: 5, File: 5}, {Rank: 6, File: 6}, {Rank: 7, File: 7},
		},
		{
			{Rank: -1, File: 1}, {Rank: -2, File: 2}, {Rank: -3, File: 3}, {Rank: -4, File: 4},
			{Rank: -5, File: 5}, {Rank: -6, File: 6}, {Rank: -7, File: 7},
		},
		{
			{Rank: 1, File: -1}, {Rank: 2, File: -2}, {Rank: 3, File: -3}, {Rank: 4, File: -4},
			{Rank: 5, File: -5}, {Rank: 6, File: -6}, {Rank: 7, File: -7},
		},
		{
			{Rank: -1, File: -1}, {Rank: -2, File: -2}, {Rank: -3, File: -3}, {Rank: -4, File: -4},
			{Rank: -5, File: -5}, {Rank: -6, File: -6}, {Rank: -7, File: -7},
		},
	}

	queenMoveOffsetLines = append(rookMoveOffsetLines, bishopMoveOffsetLines...)

	knightMoveOffsets = []Square{
		{File: -1, Rank: 2}, {File: -1, Rank: -2},
		{File: 1, Rank: 2}, {File: 1, Rank: -2},
		{File: -2, Rank: 1}, {File: -2, Rank: -1},
		{File: 2, Rank: 1}, {File: 2, Rank: -1},
	}

	pawnAdvancementOffsets = []Square{{File: 0, Rank: 1}}

	pawnInitialAdvancementOffsets = append(pawnAdvancementOffsets, Square{File: 0, Rank: 2})

	pawnCaptureOffsets = []Square{
		{File: -1, Rank: 1},
		{File: 1, Rank: 1},
	}
)
