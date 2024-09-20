package goldentour

import (
	"math"

	slot "github.com/slotopol/server/game/slot"
)

// reels lengths [32, 46, 54, 46, 32], total reshuffles 117006336
// golf bonuses: count 1264896, rtp = 17.296786%
// golf bonuses frequency: 1/92.503
// RTP = 127.57(lined) + 0(scatter) + 17.297(golf) = 144.868028%
var Reels145 = slot.Reels5x{
	{4, 6, 3, 7, 8, 2, 11, 5, 6, 4, 2, 10, 5, 8, 7, 4, 1, 7, 5, 4, 8, 7, 6, 8, 5, 6, 3, 7, 6, 5, 8, 9},
	{11, 5, 3, 8, 7, 6, 2, 8, 7, 5, 8, 4, 7, 10, 8, 4, 7, 6, 5, 7, 8, 5, 6, 8, 7, 6, 3, 7, 5, 6, 8, 7, 6, 5, 8, 7, 4, 6, 8, 2, 6, 9, 8, 1, 4, 7},
	{8, 6, 7, 8, 6, 7, 2, 6, 7, 8, 2, 7, 4, 8, 6, 4, 8, 7, 5, 6, 7, 5, 8, 1, 6, 8, 7, 5, 8, 4, 7, 6, 9, 7, 3, 4, 8, 7, 6, 8, 7, 5, 8, 6, 5, 8, 7, 5, 8, 6, 11, 8, 3, 10},
	{7, 6, 8, 7, 6, 4, 8, 3, 11, 7, 5, 6, 4, 5, 7, 6, 10, 8, 2, 5, 6, 1, 8, 6, 5, 7, 6, 8, 7, 5, 8, 7, 6, 4, 8, 7, 9, 8, 7, 3, 4, 8, 7, 5, 2, 8},
	{5, 8, 4, 5, 2, 10, 6, 4, 11, 7, 2, 6, 7, 5, 8, 4, 1, 8, 9, 6, 5, 4, 6, 3, 7, 8, 3, 7, 8, 5, 7, 6},
}

// Map with available reels.
var ReelsMap = map[float64]*slot.Reels5x{
	144.868028: &Reels145,
}

func FindReels(mrtp float64) (rtp float64, reels slot.Reels) {
	for p, r := range ReelsMap {
		if math.Abs(mrtp-p) < math.Abs(mrtp-rtp) {
			rtp, reels = p, r
		}
	}
	return
}
