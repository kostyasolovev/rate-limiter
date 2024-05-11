package cycled_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetIndex(t *testing.T) {
	t.Parallel()

	type tcase struct {
		name     string
		inp      CycledArray[int64]
		target   int64
		expIndex int
	}

	var testcases = []tcase{
		{"001",
			CycledArray[int64]{
				arr:     []int64{6, 7, 3, 4, 5},
				nextPos: 2,
				len:     5,
				cap:     5,
			},
			4,
			3},
		{"002",
			CycledArray[int64]{
				arr:     []int64{1, 0, 0, 0, 0},
				nextPos: 1,
				len:     1,
				cap:     5,
			},
			2,
			1},
		{"003",
			CycledArray[int64]{
				arr:     []int64{2, 1, 1, 1, 2},
				nextPos: 1,
				len:     5,
				cap:     5,
			},
			1,
			2},
		{"004",
			CycledArray[int64]{
				arr:     []int64{6, 8, 9, 1, 3, 4, 4, 4},
				nextPos: 3,
				len:     8,
				cap:     8,
			},
			8,
			1},
		{"005",
			CycledArray[int64]{
				arr:     []int64{1, 2, 3, 4, 5},
				nextPos: 0,
				len:     5,
				cap:     5,
			},
			3,
			2},
		{"006",
			CycledArray[int64]{
				arr:     []int64{5, 1, 2, 3, 3},
				nextPos: 1,
				len:     5,
				cap:     5,
			},
			4,
			4},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.inp.getIndex(tc.target)
			require.Equal(t, tc.expIndex, got)
		})
	}
}
